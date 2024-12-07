use std::env;

use alloy::{
    consensus::{SidecarBuilder, SimpleCoder},
    network::{EthereumWallet, TransactionBuilder, TransactionBuilder4844},
    primitives::address,
    providers::{Provider, ProviderBuilder},
    rpc::types::TransactionRequest,
    signers::local::PrivateKeySigner,
};
use eyre::Result;

#[tokio::main]
async fn main() -> Result<()> {
    let signer: PrivateKeySigner = env::var("PRIVATE_KEY")?.parse()?;
    let wallet = EthereumWallet::from(signer);
    let rpc_url = env::var("RPC_URL")?.parse()?;
    let provider = ProviderBuilder::new()
        .with_recommended_fillers()
        .wallet(wallet)
        .on_http(rpc_url);

    let sidecar: SidecarBuilder<SimpleCoder> = SidecarBuilder::from_slice(b"Hello, Ethereum!");
    let sidecar = sidecar.build()?;

    let gas_price = provider.get_gas_price().await?;
    let eip1559_est = provider.estimate_eip1559_fees(None).await?;
    let tx = TransactionRequest::default()
        .with_to(address!("d8da6bf26964af9d7eed9e03e53415d37aa96045"))
        .with_max_fee_per_blob_gas(gas_price)
        .with_max_fee_per_gas(eip1559_est.max_fee_per_gas)
        .with_max_priority_fee_per_gas(eip1559_est.max_priority_fee_per_gas)
        .with_blob_sidecar(sidecar);

    let pending_tx = provider.send_transaction(tx).await?;

    println!("Pending transaction... {}", pending_tx.tx_hash());

    let receipt = pending_tx.get_receipt().await?;

    println!(
        "Transaction included in block {}",
        receipt.block_number.expect("Failed to get block number")
    );

    Ok(())
}
