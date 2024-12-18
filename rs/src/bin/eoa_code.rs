use std::env;

use alloy::{
    eips::eip7702::Authorization,
    network::{EthereumWallet, TransactionBuilder, TransactionBuilder7702},
    providers::{Provider, ProviderBuilder},
    rpc::types::TransactionRequest,
    signers::{local::PrivateKeySigner, SignerSync},
    sol,
};
use eyre::Result;

// Codegen from embedded Solidity code and precompiled bytecode.
// solc v0.8.25 Log.sol --via-ir --optimize --bin
sol!(
    #[allow(missing_docs)]
    #[sol(rpc, bytecode = "6080806040523460135760c9908160188239f35b5f80fdfe6004361015600b575f80fd5b5f3560e01c80637b3ab2d014605f57639ee1a440146027575f80fd5b34605b575f366003190112605b577f2d67bb91f17bca05af6764ab411e86f4ddf757adb89fcec59a7d21c525d417125f80a1005b5f80fd5b34605b575f366003190112605b577fbcdfe0d5b27dd186282e187525415c57ea3077c34efb39148111e4d342e7ab0e5f80a100fea2646970667358221220f6b42b522bc9fb2b4c7d7e611c7c3e995d057ecab7fd7be4179712804c886b4f64736f6c63430008190033")]
    contract Log {
        #[derive(Debug)]
        event Hello();
        event World();

        function emitHello() public {
            emit Hello();
        }

        function emitWorld() public {
            emit World();
        }
    }
);

#[tokio::main]
async fn main() -> Result<()> {
    let signer: PrivateKeySigner = env::var("PRIVATE_KEY_1")?.parse()?;
    let authorizer: PrivateKeySigner = env::var("PRIVATE_KEY_2")?.parse()?;

    let wallet = EthereumWallet::from(signer.clone());
    let rpc_url = env::var("RPC_URL")?.parse()?;
    let provider = ProviderBuilder::new()
        .with_recommended_fillers()
        .wallet(wallet)
        .on_http(rpc_url);

    let contract = Log::deploy(&provider).await?;

    // Create an authorization object.
    let authorization = Authorization {
        chain_id: provider.get_chain_id().await?,
        // Reference to the contract that will be set as code for the authority.
        address: *contract.address(),
        nonce: provider.get_transaction_count(authorizer.address()).await?,
    };

    // Sign the authorization.
    let signature = authorizer.sign_hash_sync(&authorization.signature_hash())?;
    let signed_authorization = authorization.into_signed(signature);

    // Collect the calldata required for the transaction.
    let call = contract.emitHello();
    let emit_hello_calldata = call.calldata().to_owned();

    // Build the transaction.
    let tx = TransactionRequest::default()
        .with_to(authorizer.address())
        .with_max_fee_per_gas(100_000)
        .with_max_priority_fee_per_gas(535374991)
        .with_authorization_list(vec![signed_authorization])
        .with_input(emit_hello_calldata);

    // Send the transaction and wait for the broadcast.
    let pending_tx = provider.send_transaction(tx).await?;

    println!("Pending transaction... {}", pending_tx.tx_hash());

    // Wait for the transaction to be included and get the receipt.
    let receipt = pending_tx.get_receipt().await?;

    println!(
        "Transaction included in block {}",
        receipt.block_number.expect("Failed to get block number")
    );

    assert!(receipt.status());
    assert_eq!(receipt.from, signer.address());
    assert_eq!(receipt.to, Some(authorizer.address()));
    assert_eq!(receipt.inner.logs().len(), 1);
    assert_eq!(receipt.inner.logs()[0].address(), authorizer.address());

    Ok(())
}
