package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const password = "password"

func main() {
	b, err := os.ReadFile("wallet/UTC--2023-01-20T23-11-14.014898000Z--7f285f547f15fee627c14a44d6ad6742ff2eb645")
	if err != nil {
		log.Fatalf("Error to read the wallet file:%v", err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatalf("Error to decrypt the key:%v", err)
	}

	private_data := hexutil.Encode(crypto.FromECDSA(key.PrivateKey))
	fmt.Println("The private data from wallet:", private_data)

	public_data := hexutil.Encode(crypto.FromECDSAPub(&key.PrivateKey.PublicKey))
	fmt.Println("The public data from wallet:", public_data)

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("The address from wallet:", address)
}
