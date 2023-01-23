package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/* The balance aren't cross network */

const (
	password   = "password"
	goerli_url = "https://goerli.infura.io/v3/6a7af5f9ac4d4703812a53f49b72f75e"
)

func main() {
	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount(password)
	if err != nil {
		log.Fatalf("Error to create an account - nets: %v", err)
	}

	_, err = ks.NewAccount(password)
	if err != nil {
		log.Fatalf("Error to create an account - nets: %v", err)
	}

	client, err := ethclient.Dial(goerli_url)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	defer client.Close()

	first_address := common.HexToAddress("f9b2b8300ceda35ff834a8c05b00e471c37518f2")
	second_address := common.HexToAddress("af04853258a5d95d67d63d8c312d1a542a24b478")

	first_balance, err := client.BalanceAt(context.Background(), first_address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance from first address: %v", err)
	}

	fmt.Println("The first balance - nets:", first_balance)

	second_balance, err := client.BalanceAt(context.Background(), second_address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance from second address: %v", err)
	}

	fmt.Println("The second balance - nets:", second_balance)
}
