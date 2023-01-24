package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	goerli_websocket = "wss://goerli.infura.io/ws/v3/6a7af5f9ac4d4703812a53f49b72f75e"
	contract_address = "0xcDF8c371243420c860A5Db509B0C6eB926c68C4E"
)

func main() {
	client, err := ethclient.Dial(goerli_websocket)
	if err != nil {
		log.Fatalf("Error to create a ether client - subscribe: %v", err)
	}

	defer client.Close()

	query := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(contract_address)}}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Error to subscribe filter logs - subscribe: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}
