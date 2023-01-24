package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	wallet_filepath = "helper/wallets/UTC--2023-01-22T22-44-58.442307000Z--f9b2b8300ceda35ff834a8c05b00e471c37518f2"
	password        = "password"
)

func main() {
	b, err := os.ReadFile(wallet_filepath)
	if err != nil {
		log.Fatalf("Error to read the wallet file: %v", err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatalf("Error to decrypt the key: %v", err)
	}

	private_data := hexutil.Encode(crypto.FromECDSA(key.PrivateKey))
	fmt.Println("The private data from wallet:", private_data)

	public_data := hexutil.Encode(crypto.FromECDSAPub(&key.PrivateKey.PublicKey))
	fmt.Println("The public data from wallet:", public_data)

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("The address from wallet:", address)
}
