package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	contract "deploy/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	password   = "password"
	goerli_url = "https://goerli.infura.io/v3/6a7af5f9ac4d4703812a53f49b72f75e"
)

func main() {
	content, err := os.ReadFile("transactions/wallets/UTC--2023-01-22T22-44-58.442307000Z--f9b2b8300ceda35ff834a8c05b00e471c37518f2")
	if err != nil {
		log.Fatalf("Error to read the wallet file - deploy: %v", err)
	}

	key, err := keystore.DecryptKey(content, password)
	if err != nil {
		log.Fatalf("Error to decrypt the key - deploy: %v", err)
	}

	client, err := ethclient.Dial(goerli_url)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	defer client.Close()

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatalf("Error to get the nonce from first address - deploy: %v", err)
	}

	suggested_gas_price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error to get the suggested gas price - deploy: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error to get the chain id - deploy: %v", err)
	}

	authorization, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Error to generate a new keyed transaction with chain id: %v", err)
	}

	authorization.GasPrice = suggested_gas_price
	authorization.GasLimit = uint64(3000000)
	authorization.Nonce = big.NewInt(int64(nonce))

	adx, transaction, _, err := contract.DeployContract(authorization, client)
	if err != nil {
		log.Fatalf("Error to deploy the contract: %v", err)
	}

	fmt.Println("------------------------------")
	fmt.Println("The adx: ", adx.Hex())
	fmt.Println("The transaction: ", transaction.Hash().Hex())
	fmt.Println("------------------------------")
}
