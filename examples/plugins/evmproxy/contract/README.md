## evmproxy

This example plugin shows how to call a solidity contract running on the DAppChain's virtual machine from a diadem plugin from a web3 provider

It wraps the smart contract in SimpleStore.sol. If you have the solidity compiler solc installed
you can generate the abi and bin files with.
```bash
solc  --bin -o ./contracts SimpleStore.sol
solc  --abi -o. SimpleStore.sol
```
If you run the DAppChain in another directory, both `solc` output files need to be copied across.

## Receiving from Web3JS + diademProvider

Using Web3 + diademProvider makes possible to receive calls from Web3js from `eth_sendTransaction` and `eth_call`, under the hood
`diademProvider` will translate Web3 calls and wrap the [Contract ABI](https://solidity.readthedocs.io/en/develop/abi-spec.html) inside a diadem request which will be unpacked to call EVM inside the plugin

## Prerequisite
This example requires the 
[go-ethereum package](https://github.com/ethereum/go-ethereum),
if you do not already have it installed you can use.
```bash
go get -d github.com/ethereum/go-ethereum
```

## Install
Amend the location entry in the `genesis.json` for SimpleStore to match the 
path of your go-diadem directory.

`ed genesis.json`

Build the evmproxy library.
Assuming diadem is in your path initialise the DAppChain `./diadem init`. 
copy the `example.genesis.json` to `genesis.json` and then 
run the chain with `./diadem run`

```bash
go build -tags "evm" -o contracts/evmproxy.1.0.0  evmproxy.go
diadem init
cp example.genesis.json genesis.json
diadem run
```

## Testing
The cli can be bult with
```bash
cd ../cli
go build -o evmproxy-cli evmproxy.go
```

You can now run the cli/evmproxy.go tool to access the solidty contract.
You might need to use -r and -w to set the DAppChain's URL.
```bash
./evmproxy-cli set
./evmproxy-cli get
```

