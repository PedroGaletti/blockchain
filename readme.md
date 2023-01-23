## Blockchain

**Warning**: I dedicated this repository to my Blockchain study with Golang, but if you want to contribute with some content, feel free.

## Stack

- [Golang](https://go.dev) - Build fast, reliable, and efficient software at scale

## Content

- [Balance](https://github.com/PedroGaletti/blockchain/tree/main/balance/main.go) - Get the balance from transaction
- [Deploy](https://github.com/PedroGaletti/blockchain/tree/main/deploy/main.go) - Deploy the smart contract
- [Contract](https://github.com/PedroGaletti/blockchain/tree/main/deploy/contracts/contract.sol) - Example of smart contract
- [Keys](https://github.com/PedroGaletti/blockchain/tree/main/keys/main.go) - Generate key (private, public, address)
- [Nets](https://github.com/PedroGaletti/blockchain/tree/main/nets/main.go) - Working with networks in the blockchain
- [Store](https://github.com/PedroGaletti/blockchain/tree/main/store/main.go) - Store the secret file of key
- [Transactions](https://github.com/PedroGaletti/blockchain/tree/main/transactions/main.go) - How to make transactions
- [Wallet](https://github.com/PedroGaletti/blockchain/tree/main/wallet/main.go) - Get the key by the wallet security file

## Commands

With the `solc` command you can generate the .bin and .abi files of your contract:
```
solc --bin --abi deploy/contracts/contract.sol -o build
```

With the `abigen` you can generate the .go file that will make all the functions to you deploy your contract:
```
abigen --bin=build/Contract.bin --abi=build/Contract.abi --pkg=contract --out=deploy/contracts/contract.go
```

## Usefull Websites

- [Solidity](https://docs.soliditylang.org) - This is the Solidity website, you can install the solidity compiler in the link.
- [Goerli Ether Scan](https://goerli.etherscan.io) - This is a website to search the address of contracts and transactions.
- [Goerli Faucet](https://goerlifaucet.com) - On this website, you can add some Ethereum to a wallet.
- [Infura](https://www.infura.io/) - Infura it's a platform to you building your blockchain application.
- [Ganache](https://www.npmjs.com/package/ganache) - Ganache is an Ethereum simulator that makes developing Ethereum applications.
- [Protoc](https://github.com/protocolbuffers/protobuf/releases) - Link to install the protoc in your desktop.

## How to use

I'm using golang 1.19, and since golang 1.18 we can work with **go.work**, so to use and see the things that I implemented, it's necessary to run the following command:

```
go run ${module that you want}
```

Like:

```
go run keys
```