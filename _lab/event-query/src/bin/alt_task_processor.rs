use alloy::{
    primitives::{keccak256, Address, FixedBytes},
    providers::{Provider, ProviderBuilder, RootProvider},
    rpc::types::Filter,
    sol,
    sol_types::SolEvent,
    transports::http::{Client, Http},
};
use event_query::constants::{BLOCK_RANGE, CONTRACT_ADDRESS, EXAM_TITLE, RPC_URL};
use futures::future::join_all;
use std::error::Error;
use std::sync::Arc;
use tokio::sync::{Mutex, Semaphore};
use tokio::{self, spawn};
use Datastore::Stored;

sol!(
    #[sol(rpc)]
    Datastore,
    "common/Datastore.json"
);

async fn process_block_range(
    provider: &RootProvider<Http<Client>>,
    filter_topic: FixedBytes<32>,
    start: u64,
    end: u64,
) -> Result<Vec<Stored>, Box<dyn Error + Send + Sync>> {
    let mut events: Vec<Stored> = Vec::new();

    let filter = Filter::new()
        .address(CONTRACT_ADDRESS.parse::<Address>()?)
        .event(Datastore::Stored::SIGNATURE)
        .from_block(start)
        .to_block(end);
    let logs = provider.get_logs(&filter).await?;
    for l in logs {
        if l.topics()[1] == filter_topic {
            let parsed = Datastore::Stored::decode_log_data(l.data(), true)?;
            events.push(parsed);
        }
    }
    Ok(events)
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

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error + Send + Sync>> {
    let max_concurrent = 7;

    // Create shared vector to store all logs
    let all_logs = Arc::new(Mutex::new(Vec::new()));

    // Create Semaphore to limit concurrent tasks
    let semaphore = Arc::new(Semaphore::new(max_concurrent));

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

    let block_ranges =
        calculate_block_ranges(range.start.to::<u64>(), range.end.to::<u64>(), BLOCK_RANGE);

    // Create tasks for each chunk
    let tasks = block_ranges.into_iter().map(|(start, end)| {
        let provider_clone = provider.clone();
        let logs_clone = Arc::clone(&all_logs);
        let sem_clone = Arc::clone(&semaphore);

        spawn(async move {
            // Acquire semaphore permit
            let permit = sem_clone.acquire().await.unwrap();

            let result = process_block_range(&provider_clone, filter_topic, start, end).await;
            match result {
                Ok(logs) => {
                    let mut all_logs = logs_clone.lock().await;
                    all_logs.extend(logs);
                    println!("Processed blocks {} to {}", start, end);
                }
                Err(e) => eprintln!("Error processing blocks {} to {}: {}", start, end, e),
            }

            // Permit is automatically dropped here, releasing the semaphore slot
            drop(permit);
        })
    });

    // Wait for all tasks to complete
    let results = join_all(tasks).await;

    // Check for any errors in task execution
    for result in results {
        if let Err(e) = result {
            eprintln!("Task panicked: {}", e);
        }
    }

    // Access final results
    let final_logs = all_logs.lock().await;
    println!("Total logs collected: {}", final_logs.len());

    Ok(())
}
