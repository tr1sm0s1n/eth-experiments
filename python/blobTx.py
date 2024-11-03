import os
from eth_abi import abi
from eth_account import Account
from eth_utils import to_hex
from web3 import Web3, HTTPProvider


rpc_url = os.environ.get("RPC_URL")
w3 = Web3(HTTPProvider(rpc_url))

text = "<( o.O )>"
encoded_text = abi.encode(["string"], [text])

BLOB_DATA = (b"\x00" * 32 * (4096 - len(encoded_text) // 32)) + encoded_text

pkey = os.environ.get("PRIVATE_KEY")
acct = w3.eth.account.from_key(pkey)

tx = {
    "type": 3,
    "chainId": w3.eth.chain_id,
    "from": acct.address,
    "to": "0x0000000000000000000000000000000000000000",
    "value": 0,
    "maxFeePerGas": 10**12,
    "maxPriorityFeePerGas": 10**12,
    "maxFeePerBlobGas": to_hex(10**12),
    "nonce": w3.eth.get_transaction_count(acct.address),
}

gas_estimate = w3.eth.estimate_gas(tx)
tx["gas"] = gas_estimate

signed = acct.sign_transaction(tx, blobs=[BLOB_DATA])
tx_hash = w3.eth.send_raw_transaction(signed.rawTransaction)
tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Tx receipt: {tx_receipt}")
