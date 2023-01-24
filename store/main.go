package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const password = "password"

func main() {
	key := keystore.NewKeyStore("helper/wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := key.NewAccount(password)
	if err != nil {
		log.Fatalf("Error to create the account: %v", err)
	}

	fmt.Println("The account address:", account.Address)
}
