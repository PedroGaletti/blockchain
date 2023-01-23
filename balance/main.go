package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	infura_uri  = "https://mainnet.infura.io/v3/6a7af5f9ac4d4703812a53f49b72f75e"
	ganache_uri = "http://127.0.0.1:8545"
)

func main() {
	client, err := ethclient.DialContext(context.Background(), infura_uri)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}

	fmt.Println("The block number:", block.Number())

	address := common.HexToAddress("0x53f2Dd7Ab552bFffc4d264B69f35AEC597bd78c3")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance: %v", err)
	}

	fmt.Println("The balance:", balance) // 1 ether == 10^18 wei

	float_balance := new(big.Float)
	float_balance.SetString(balance.String())
	fmt.Println("The float balance:", float_balance)

	ether_balance := new(big.Float).Quo(float_balance, big.NewFloat(math.Pow10(18))) // 10 * 10 * 10 * 10 ......
	fmt.Println("The ether balance:", ether_balance)
}
