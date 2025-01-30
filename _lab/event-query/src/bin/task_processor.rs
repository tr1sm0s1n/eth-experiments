use std::sync::Arc;

use alloy::{
    primitives::{keccak256, Address, FixedBytes},
    providers::{Provider, ProviderBuilder, RootProvider},
    rpc::types::Filter,
    sol,
    sol_types::SolEvent,
    transports::http::{Client, Http},
};
use event_query::constants::{BLOCK_RANGE, CONTRACT_ADDRESS, EXAM_TITLE, RPC_URL};
use eyre::Result;
use futures::{stream::FuturesUnordered, StreamExt};
use tokio::{
    spawn,
    sync::{Mutex, Semaphore},
};
use Datastore::Stored;

sol!(
    #[sol(rpc)]
    Datastore,
    "common/Datastore.json"
);

#[tokio::main]
async fn main() -> Result<()> {
    // Use a concurrent-friendly data structure for better performance
    let events: Arc<Mutex<Vec<Stored>>> = Arc::new(Mutex::new(Vec::new()));
    let semaphore = Arc::new(Semaphore::new(6));
    let mut fetch_futures = FuturesUnordered::new();

    let rpc_url = RPC_URL.parse()?;
    let provider = ProviderBuilder::new().on_http(rpc_url);
    let instance = Datastore::new(CONTRACT_ADDRESS.parse()?, provider.clone());

    // Fetch range first to avoid repeated calls
    let range = instance.EventCount(EXAM_TITLE.to_string()).call().await?;
    let filter_topic = keccak256(EXAM_TITLE);
    println!(
        "Data Range: [\x1b[1;36m{:?}\x1b[0m] -> [\x1b[1;36m{:?}\x1b[0m]",
        range.start.to::<u64>(),
        range.end.to::<u64>()
    );

    // Calculate block ranges upfront to reduce runtime calculations
    let block_ranges =
        calculate_block_ranges(range.start.to::<u64>(), range.end.to::<u64>(), BLOCK_RANGE);

    // Use join_all for more efficient concurrent processing
    for (start, end) in block_ranges {
        let provider_clone = provider.clone();
        let events_clone = events.clone();
        let semaphore_clone = Arc::clone(&semaphore);

        fetch_futures.push(spawn(async move {
            let _permit = semaphore_clone.acquire().await.unwrap(); // Wait for a slot
            fetch_logs(&provider_clone, filter_topic, &events_clone, start, end).await
        }));
    }

    // Wait for all fetch tasks to complete
    while let Some(result) = fetch_futures.next().await {
        // Handle any errors from the fetching tasks
        if let Err(err) = result {
            eprintln!("Task failed: {:?}", err);
        }
    }

    // Atomic read of events length for thread-safe logging
    let event_count = events.lock().await.len();
    println!("Processed \x1b[1;34m{event_count}\x1b[0m event logs!!");

    Ok(())
}

async fn fetch_logs(
    provider: &RootProvider<Http<Client>>,
    filter_topic: FixedBytes<32>,
    events: &Arc<Mutex<Vec<Stored>>>,
    start: u64,
    end: u64,
) -> Result<()> {
    println!(
        "Processing: [\x1b[1;32m{:?}\x1b[0m] -> [\x1b[1;31m{:?}\x1b[0m]",
        start, end
    );
    let filter = Filter::new()
        .address(CONTRACT_ADDRESS.parse::<Address>()?)
        .event(Datastore::Stored::SIGNATURE)
        .from_block(start)
        .to_block(end);

    let logs = provider.get_logs(&filter).await?;
    for l in logs {
        if l.topics()[1] == filter_topic {
            let parsed = Datastore::Stored::decode_log_data(l.data(), true)?;
            let mut events_guard = events.lock().await;
            events_guard.push(parsed);
        }
    }

    Ok(())
}

// Helper function to generate block ranges
fn calculate_block_ranges(start: u64, end: u64, block_range: u64) -> Vec<(u64, u64)> {
    let mut ranges = Vec::new();
    let mut current_start = start;

    while current_start <= end {
        let current_end = (current_start + block_range).min(end);
        ranges.push((current_start, current_end));
        current_start = current_end + 1;
    }

    ranges
}
