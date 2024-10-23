# ETH-Experiments

Reference repository for experimenting with Go-Ethereum.

## 🛠 Built With

[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)
[![Geth Badge](https://img.shields.io/badge/Geth-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://geth.ethereum.org/)

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
