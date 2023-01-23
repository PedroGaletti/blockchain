package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	private_key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Error to generate a private key: %v", err)
	}

	private_data := crypto.FromECDSA(private_key)
	fmt.Println("The private data:", private_data)
	exactly_private_key := hexutil.Encode(private_data)
	fmt.Println("The exactly private key:", exactly_private_key)

	public_data := crypto.FromECDSAPub(&private_key.PublicKey)
	fmt.Println("The public data:", public_data)
	exactly_public_key := hexutil.Encode(public_data)
	fmt.Println("The exactly public key:", exactly_public_key)

	address := crypto.PubkeyToAddress(private_key.PublicKey).Hex()
	fmt.Println("The public address:", address)
}
