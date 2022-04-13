# Rosetta-DBC

coinbase rosetta dbc

```sh
go build

# Testnet
MODE=ONLINE NETWORK=TESTNET PORT="8080" ./rosetta-dbc run
rosetta-cli check:data --configuration-file rosetta-cli-conf/testnet/config.json
rosetta-cli check:construction --configuration-file rosetta-cli-conf/testnet/config.json

# Mainnet
MODE=ONLINE NETWORK=MAINNET PORT="8080" ./rosetta-dbc run
cd
rosetta-cli check:data --configuration-file rosetta-cli-conf/mainnet/config.json
rosetta-cli check:construction --configuration-file rosetta-cli-conf/mainnet/config.json
```

> Related repos:

https://github.com/centrifuge/go-substrate-rpc-client.git

https://github.com/coinbase/rosetta-sdk-go.git

https://github.com/docknetwork/rosetta-api.git

https://github.com/coinbase/rosetta-ethereum.git

https://github.com/coinbase/rosetta-cli.git
