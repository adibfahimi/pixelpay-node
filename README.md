# Pixelpay Node

Pixelpay is a simple blockchain project.

## Start the Pixelpay Node

You have two options to start the Pixelpay Node:

### Option 1: Local

Download the executable file from [releases](https://github.com/adibfahimi/pixelpay-node/releases):

```bash
./pixelpay-node
```

### Option 2: Using Docker

```bash
git clone https://github.com/adibfahimi/pixelpay-node
cd pixelpay-node
docker build -t pixelpay-node .
docker run -p 3000:3000 pixelpay-node
```

This will start the Pixelpay Go server on `http://localhost:3000`.

## Related Projects

- [Pixelpay Wallet](https://github.com/adibfahimi/pixelpay-wallet)
- [Pixelpay Miner](https://github.com/adibfahimi/pixelpay-miner)
