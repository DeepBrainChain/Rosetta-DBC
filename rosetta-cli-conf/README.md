```
wget https://github.com/DeepBrainChain/DeepBrainChain-MainChain/releases/download/v2/dbc_chain_linux_x64.tar.gz -O dbc_chain_linux_x64.tar.gz
tar xf dbc_chain_linux_x64.tar.gz && cd dbc-chain-mainnet
```

For development node:

```
rosetta-cli --configuration-file=testnet/config.json check:construction
rosetta-cli --configuration-file=testnet/config.json check:data
```

For Dock mainnet:

```
rosetta-cli --configuration-file=mainnet/config.json check:xxx
```
