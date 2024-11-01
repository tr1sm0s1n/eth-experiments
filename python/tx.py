from web3 import Web3, EthereumTesterProvider

def main():
    w3 = Web3(EthereumTesterProvider())

    # eth-tester populates accounts with test ether:
    acct1 = w3.eth.accounts[0]

    some_address = "0x0000000000000000000000000000000000000000"

    # when using one of its generated test accounts,
    # eth-tester signs the tx (under the hood) before sending:
    tx_hash = w3.eth.send_transaction({
        "from": acct1,
        "to": some_address,
        "value": 123123123123123
    })

    tx = w3.eth.get_transaction(tx_hash)
    assert tx["from"] == acct1 


if __name__ == "__main__":
    main()
