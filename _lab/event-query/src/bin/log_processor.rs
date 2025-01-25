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
use Datastore::Stored;

sol!(
    #[sol(rpc)]
    Datastore,
    "common/Datastore.json"
);

#[tokio::main]
async fn main() -> Result<()> {
    // To store the parsed events.
    let mut events: Vec<Stored> = vec![];

    // Create the provider.
    let rpc_url = RPC_URL.parse()?;
    let provider = ProviderBuilder::new().on_http(rpc_url);
    // Create the instance.
    let instance = Datastore::new(CONTRACT_ADDRESS.parse()?, provider.clone());
    // Fetch the block range.
    let range = instance.EventCount(EXAM_TITLE.to_string()).call().await?;

    let filter_topic = keccak256(EXAM_TITLE);

    let mut start = range.start.to::<u64>();
    let mut end = range.start.to::<u64>() + BLOCK_RANGE;

    loop {
        if end >= range.end.to::<u64>() {
            end = range.end.to::<u64>();
            fetch_logs(&provider, filter_topic, &mut events, start, end).await?;
            break;
        }
        fetch_logs(&provider, filter_topic, &mut events, start, end).await?;
        start = end + 1;
        end = end + 1 + BLOCK_RANGE;
    }

    println!("Processed \x1b[34m{:?}\x1b[0m event logs!!", events.len());

    Ok(())
}

async fn fetch_logs(
    provider: &RootProvider<Http<Client>>,
    filter_topic: FixedBytes<32>,
    events: &mut Vec<Stored>,
    start: u64,
    end: u64,
) -> Result<()> {
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

    Ok(())
}
