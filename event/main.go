package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	contract "deploy/contracts"
)

const (
	goerli_websocket = "wss://goerli.infura.io/ws/v3/6a7af5f9ac4d4703812a53f49b72f75e"
	contract_address = "0xcDF8c371243420c860A5Db509B0C6eB926c68C4E"
)

func main() {
	client, err := ethclient.Dial(goerli_websocket)
	if err != nil {
		log.Fatalf("Error to create a ether client - event: %v", err)
	}

	query := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(contract_address)}}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Error to filter logs - event: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractABI)))
	if err != nil {
		log.Fatalf("Error to get the abi json - event: %v", err)
	}

	for _, vLog := range logs {
		fmt.Println("vLog.BlockHash.Hex(): ", vLog.BlockHash.Hex())
		fmt.Println("vLog.BlockNumber: ", vLog.BlockNumber)
		fmt.Println("vLog.TxHash.Hex(): ", vLog.TxHash.Hex())

		i, err := contractAbi.Unpack("LotteryEvent", vLog.Data)
		if err != nil {
			log.Fatalf("Error to get the contract abi unpack - event: %v", err)
		}

		fmt.Println("interface:", i)

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println("first topic: ", topics[0])
	}

	eventSignature := []byte("GuessNumber(uint256)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("Hash hex: ", hash.Hex())
}
