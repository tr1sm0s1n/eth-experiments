from web3 import Web3
from os import getenv
from dotenv import load_dotenv
from eth_abi import encode

load_dotenv()
w3 = Web3(Web3.HTTPProvider(getenv("RPC_URL")))
eoa = w3.eth.account.from_key(getenv("PRIVATE_KEY"))  # replace with your own set up

tx = {
    "from": eoa.address,
    "value": 0,
    "chainId": w3.eth.chain_id,
    "gas": 250000,
    "to": "0x0000000000000000000000000000000000000000",
    "maxFeePerGas": w3.eth.gas_price * 2,
    "maxPriorityFeePerGas": w3.eth.max_priority_fee * 2,
    "maxFeePerBlobGas": hex(10**12),
    "nonce": w3.eth.get_transaction_count(eoa.address),
}

text = "<( o.O )>"
encoded_text = encode(["string"], [text])

padding_bytes = b"\x00" * 32 * (4096 - len(encoded_text) // 32)

BLOB_DATA = padding_bytes + encoded_text

signed_tx = eoa.sign_transaction(tx, blobs=[BLOB_DATA])
tx_hash = w3.eth.send_raw_transaction(signed_tx.raw_transaction)
print(f"Tx hash: 0x{tx_hash.hex()}")
tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Tx receipt: {tx_receipt}")
