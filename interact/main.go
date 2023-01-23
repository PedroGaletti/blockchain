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
	"github.com/ethereum/go-ethereum/common"
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
		log.Fatalf("Error to read the wallet file - interact: %v", err)
	}

	key, err := keystore.DecryptKey(content, password)
	if err != nil {
		log.Fatalf("Error to decrypt the key - interact: %v", err)
	}

	client, err := ethclient.Dial(goerli_url)
	if err != nil {
		log.Fatalf("Error to create a ether client - interact: %v", err)
	}

	defer client.Close()

	suggested_gas_price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error to get the suggested gas price - interact: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error to get the chain id - interact: %v", err)
	}

	contract_address := common.HexToAddress("0xD52bBaCfCa6D61FeC9451268d8722d5D645bF380") // contract address
	new_contract, err := contract.NewContract(contract_address, client)
	if err != nil {
		log.Fatalf("Error to create a new contract - interact: %v", err)
	}

	options, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatalf("Error to create a new keyed transaction with chain id - interact: %v", err)
	}
	options.GasLimit = 3000000
	options.GasPrice = suggested_gas_price

	transaction, err := new_contract.Add(options, "First task")
	if err != nil {
		log.Fatalf("Error to add options in the new contract - interact: %v", err)
	}

	fmt.Println(transaction.Hash())

	key_address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	tasks, err := new_contract.List(&bind.CallOpts{From: key_address})
	if err != nil {
		log.Fatalf("Error to list tasks from the new contract - interact: %v", err)
	}

	fmt.Println("Tasks: ", tasks)

	updated_transaction, err := new_contract.Update(options, big.NewInt(0), "Update first task")
	if err != nil {
		log.Fatalf("Error to update task from the new contract - interact: %v", err)
	}

	fmt.Println("Updated transaction", updated_transaction.Hash())

	toggled_transaction, err := new_contract.Toggle(options, big.NewInt(0))
	if err != nil {
		log.Fatalf("Error to toggle task from the new contract - interact: %v", err)
	}

	fmt.Println("Toggled transaction", toggled_transaction.Hash())

	removed_transaction, err := new_contract.Remove(options, big.NewInt(0))
	if err != nil {
		log.Fatalf("Error to remove task from the new contract - interact: %v", err)
	}

	fmt.Println("Removed transaction", removed_transaction.Hash())
}
