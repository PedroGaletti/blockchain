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
	password        = "password"
	goerli_url      = "wss://ws.testnet.cloudwalk.io"
	wallet_filepath = "helper/wallets/UTC--2023-01-22T22-44-59.367448000Z--af04853258a5d95d67d63d8c312d1a542a24b478"
)

func main() {
	content, err := os.ReadFile(wallet_filepath)
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
