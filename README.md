# ETH-Experiments

Reference repository for experimenting with Go-Ethereum and web3.py.

## 🛠 Built With

[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)
[![Geth Badge](https://img.shields.io/badge/Geth-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://geth.ethereum.org/)
[![Python Badge](https://img.shields.io/badge/Python-3776AB?logo=python&logoColor=fff&style=for-the-badge)](https://www.python.org/)
[![web3.py Badge](https://img.shields.io/badge/web3.py-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://web3py.readthedocs.io/en/stable/)

## Blob Transaction Example

Create a `.env` file and add the following:

```sh
RPC_URL=TBD
PRIVATE_KEY=TBD
```

Send a blob transaction to the given network:

```sh
go run public/blobTx.go
```
