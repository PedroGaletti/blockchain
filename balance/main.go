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

const goerli_url = "wss://goerli.infura.io/ws/v3/6a7af5f9ac4d4703812a53f49b72f75e"

func main() {
	client, err := ethclient.DialContext(context.Background(), goerli_url)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}

	fmt.Println("The block number:", block.Number())

	address := common.HexToAddress("f9b2b8300ceda35ff834a8c05b00e471c37518f2")
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
