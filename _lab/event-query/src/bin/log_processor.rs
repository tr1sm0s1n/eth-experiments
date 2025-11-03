use std::cmp::min;

use alloy::{
    primitives::{keccak256, Address, FixedBytes},
    providers::{Provider, ProviderBuilder},
    rpc::types::Filter,
    sol,
    sol_types::SolEvent,
};
use event_query::{
    constants::{BLOCK_RANGE, CONTRACT_ADDRESS, EXAM_TITLE, RPC_URL},
    types::MyProvider,
};
use eyre::Result;
use DataStore::Stored;

sol!(
    #[sol(rpc)]
    DataStore,
    "artifacts/DataStore.json"
);

#[tokio::main]
async fn main() -> Result<()> {
    // To store the parsed events.
    let mut events: Vec<Stored> = vec![];

    // Create the provider.
    let rpc_url = RPC_URL.parse()?;
    let provider: MyProvider = ProviderBuilder::new().connect_http(rpc_url);
    // Create the instance.
    let instance = DataStore::new(CONTRACT_ADDRESS.parse()?, provider.clone());
    // Fetch the block range.
    let range = instance.EventCount(EXAM_TITLE.to_string()).call().await?;
    println!(
        "Data Range: [\x1b[1;36m{:?}\x1b[0m] -> [\x1b[1;36m{:?}\x1b[0m]",
        range.start.to::<u64>(),
        range.end.to::<u64>()
    );

    let filter_topic = keccak256(EXAM_TITLE);

    let mut start = range.start.to::<u64>();
    let mut end;

    loop {
        if start > range.end.to::<u64>() {
            break;
        }
        end = min(start + BLOCK_RANGE, range.end.to::<u64>());
        println!(
            "Processing: [\x1b[1;32m{:?}\x1b[0m] -> [\x1b[1;31m{:?}\x1b[0m]",
            start, end
        );
        fetch_logs(&provider, filter_topic, &mut events, start, end).await?;
        start = end + 1;
    }

    println!("Processed \x1b[1;34m{:?}\x1b[0m event logs!!", events.len());

    Ok(())
}

async fn fetch_logs(
    provider: &MyProvider,
    filter_topic: FixedBytes<32>,
    events: &mut Vec<Stored>,
    start: u64,
    end: u64,
) -> Result<()> {
    let filter = Filter::new()
        .address(CONTRACT_ADDRESS.parse::<Address>()?)
        .event(DataStore::Stored::SIGNATURE)
        .from_block(start)
        .to_block(end);

    let logs = provider.get_logs(&filter).await?;
    for l in logs {
        if l.topics()[1] == filter_topic {
            let parsed = DataStore::Stored::decode_log_data(l.data())?;
            events.push(parsed);
        }
    }

    Ok(())
}
