# web3.py

## Blob Transaction Example

Create a `.env` file and add the following:

```sh
RPC_URL=TBD
PRIVATE_KEY=TBD
```

Install `uv`, an extremely fast Python package and project manager:

```sh
curl -LsSf https://astral.sh/uv/install.sh | sh
```

Send a blob transaction to the given network:

```sh
uv run public/blobTx.py
```
