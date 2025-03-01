use alloy::{
    primitives::Address,
    providers::{Provider, ProviderBuilder},
    rpc::{
        client::WsConnect,
        types::eth::{BlockNumberOrTag, Filter},
    },
    sol,
    sol_types::SolEvent,
};
use event_query::constants::{CONTRACT_ADDRESS, WS_URL};
use eyre::Result;
use futures::stream::StreamExt;

sol!(
    #[sol(rpc)]
    Datastore,
    "common/Datastore.json"
);

#[tokio::main]
async fn main() -> Result<()> {
    // Create the provider.
    let ws = WsConnect::new(WS_URL);
    let provider = ProviderBuilder::new().on_ws(ws).await?;

    let filter = Filter::new()
        .address(CONTRACT_ADDRESS.parse::<Address>().unwrap())
        // By specifying an `event` or `event_signature` we listen for a specific event of the contract.
        .event(Datastore::Stored::SIGNATURE)
        .from_block(BlockNumberOrTag::Latest);

    // Subscribe to logs.
    let sub = provider.subscribe_logs(&filter).await?;
    let mut stream = sub.into_stream();

    println!("Listening for events...");
    println!("-----------------------");

    while let Some(log) = stream.next().await {
        let parsed_log = Datastore::Stored::decode_log_data(log.data(), true).unwrap();
        println!("Event occured!!");
        println!("--------------------");
        println!("Exam No: \x1b[36m{:?}\x1b[0m", parsed_log.exam_no);
        println!("Data: \x1b[36m{:?}\x1b[0m", parsed_log.data);
        println!("--------------------");
    }

    Ok(())
}
