use alloy::{
    primitives::{keccak256, Address},
    providers::{Provider, ProviderBuilder},
    rpc::types::Filter,
    sol,
    sol_types::SolEvent,
};
use event_query::constants::{CONTRACT_ADDRESS, EXAM_TITLE, RPC_URL};
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

    let filter = Filter::new()
        .address(CONTRACT_ADDRESS.parse::<Address>()?)
        .event(Datastore::Stored::SIGNATURE)
        .from_block(range.start.to::<u64>())
        .to_block(range.end.to::<u64>());

    let logs = provider.get_logs(&filter).await?;
    let filter_topic = keccak256(EXAM_TITLE);

    for l in logs {
        if l.topics()[1] == filter_topic {
            let parsed = Datastore::Stored::decode_log_data(l.data(), true)?;
            events.push(parsed);
        }
    }

    println!("Processed \x1b[34m{:?}\x1b[0m event logs!!", events.len());

    Ok(())
}
