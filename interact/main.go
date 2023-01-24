package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"
	"time"

	contract "deploy/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	password                  = "password"
	goerli_websocket          = "wss://goerli.infura.io/ws/v3/6a7af5f9ac4d4703812a53f49b72f75e"
	first_wallet_filepath     = "helper/wallets/UTC--2023-01-22T22-44-58.442307000Z--f9b2b8300ceda35ff834a8c05b00e471c37518f2"
	second_wallet_filepath    = "helper/wallets/UTC--2023-01-22T22-44-59.367448000Z--af04853258a5d95d67d63d8c312d1a542a24b478"
	contract_address          = "0xcDF8c371243420c860A5Db509B0C6eB926c68C4E"
	transactions_stress_times = 25
)

func decodeWallets(key_filepath string, client *ethclient.Client) Wallet {
	wallet, err := os.ReadFile(key_filepath)
	if err != nil {
		log.Fatalf("Error to read the wallet file - interact: %v", err)
	}

	key, err := keystore.DecryptKey(wallet, password)
	if err != nil {
		log.Fatalf("Error to decrypt the key - interact: %v", err)
	}

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatalf("Error to get the nonce from first_key - deploy: %v", err)
	}

	return Wallet{Wallet: wallet, Key: key, Address: address, Nonce: nonce}
}

func main() {
	var wg sync.WaitGroup

	client, err := ethclient.Dial(goerli_websocket)
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

	new_contract, err := contract.NewContract(common.HexToAddress(contract_address), client)
	if err != nil {
		log.Fatalf("Error to create a new contract - interact: %v", err)
	}

	wg.Add(2)
	go addTransactionToWallet(&wg, decodeWallets(first_wallet_filepath, client), chainID, suggested_gas_price, new_contract, client)
	go addTransactionToWallet(&wg, decodeWallets(second_wallet_filepath, client), chainID, suggested_gas_price, new_contract, client)
	wg.Wait()
}

func addTransactionToWallet(wg *sync.WaitGroup, w Wallet, chainID *big.Int, suggested_gas_price *big.Int, new_contract *contract.Contract, client *ethclient.Client) {
	var (
		x            int64
		transactions []common.Hash
	)

	for x = 0; x < transactions_stress_times; x++ {
		options, err := bind.NewKeyedTransactorWithChainID(w.Key.PrivateKey, chainID)
		if err != nil {
			log.Fatalf("Error to create a new keyed transaction with chain id wallet: %s - interact: %v", w.Address.String(), err)
		}
		options.GasLimit = 3000000
		options.GasPrice = suggested_gas_price
		options.Nonce = big.NewInt(int64(int64(w.Nonce) + x))

		transaction, err := new_contract.GuessNumber(options, big.NewInt(x))
		if err != nil {
			log.Fatalf("Error to add options in the new contract wallet address: %s - interact: %v", w.Address.String(), err)
		}

		fmt.Printf("this is the transaction wallet address: %s / %s \n", w.Address.String(), transaction.Hash())
		transactions = append(transactions, transaction.Hash())
	}

	time.Sleep(15 * time.Second)

	for _, tx := range transactions {
		receipt, err := client.TransactionReceipt(context.Background(), tx)
		if err != nil {
			log.Fatalf("Error to get receipt wallet address: %s - interact: %v", w.Address.String(), err)
		}

		fmt.Printf("wallet address: %s / transaction hash: %s and status: %d \n", w.Address.String(), tx.String(), receipt.Status)
	}

	wg.Done()
}
