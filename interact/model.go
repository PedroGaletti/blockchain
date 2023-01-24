package main

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

type Wallet struct {
	Wallet  []byte
	Key     *keystore.Key
	Address common.Address
	Nonce   uint64
}
