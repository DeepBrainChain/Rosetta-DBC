# Rosetta-DBC

coinbase rosetta dbc

```sh
# Testcase
go build && MODE=ONLINE NETWORK=TESTNET PORT="8080" ./rosetta-dbc run
rosetta-cli --configuration-file=testnet/config.json check:data

# Mainnet
MODE=ONLINE NETWORK=MAINNET PORT="8080" ./rosetta-dbc run
rosetta-cli --configuration-file=mainnet/config.json check:data
```

> Related repos:

https://github.com/centrifuge/go-substrate-rpc-client.git

https://github.com/coinbase/rosetta-sdk-go.git

https://github.com/docknetwork/rosetta-api.git

https://github.com/coinbase/rosetta-ethereum.git
