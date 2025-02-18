# Go-Ethereum

## Blob Transaction Example

Create a `.env` file and add the following:

```sh
RPC_URL=TBD
PRIVATE_KEY=TBD
```

Send a blob transaction to the given network:

```sh
go run public/blobs/blobTx.go
```

## Set EOA Account Code Transaction Example

Add multiple private keys to the `.env` file:

```sh
PRIVATE_KEY_1=TBD
PRIVATE_KEY_2=TBD
```

Send a EOA account code transaction to the given network:

```sh
go run public/EOA/setCodeTx.go
```
