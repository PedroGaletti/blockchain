package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	password   = "password"
	goerli_url = "https://goerli.infura.io/v3/6a7af5f9ac4d4703812a53f49b72f75e"
)

func main() {
	client, err := ethclient.Dial(goerli_url)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	defer client.Close()

	first_address := common.HexToAddress("f9b2b8300ceda35ff834a8c05b00e471c37518f2")  // address from: transactions/wallets/UTC--2023-01-22T22-44-58.442307000Z--f9b2b8300ceda35ff834a8c05b00e471c37518f2
	second_address := common.HexToAddress("af04853258a5d95d67d63d8c312d1a542a24b478") // address from: transactions/wallets/UTC--2023-01-22T22-44-59.367448000Z--af04853258a5d95d67d63d8c312d1a542a24b478

	first_balance, err := client.BalanceAt(context.Background(), first_address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance from first address: %v", err)
	}

	fmt.Println("The first balance - transactions:", first_balance)

	second_balance, err := client.BalanceAt(context.Background(), second_address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance from second address: %v", err)
	}

	fmt.Println("The second balance - transactions:", second_balance)

	first_nonce, err := client.PendingNonceAt(context.Background(), first_address)
	if err != nil {
		log.Fatalf("Error to get the nonce from first address - transactions: %v", err)
	}

	amount := big.NewInt(100) // 1 ether = 1,000,000,000,000,000,000  // 10 ^ 18 // 1000000000000000000
	suggested_gas_price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error to get the suggested gas price: %v", err)
	}

	unsigned_transaction := types.NewTransaction(first_nonce, second_address, amount, 21000, suggested_gas_price, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error to get the chain id: %v", err)
	}

	content, err := os.ReadFile("transactions/wallets/UTC--2023-01-22T22-44-58.442307000Z--f9b2b8300ceda35ff834a8c05b00e471c37518f2")
	if err != nil {
		log.Fatalf("Error to read the wallet file - transactions: %v", err)
	}

	key, err := keystore.DecryptKey(content, password)
	if err != nil {
		log.Fatalf("Error to decrypt the key - transactions: %v", err)
	}

	signed_transaction, err := types.SignTx(unsigned_transaction, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatalf("Error to sign the transaction: %v", err)
	}

	if err := client.SendTransaction(context.Background(), signed_transaction); err != nil {
		log.Fatalf("Error to send the transaction: %v", err)
	}

	fmt.Println("Transaction sent: ", signed_transaction.Hash().Hex())
}
