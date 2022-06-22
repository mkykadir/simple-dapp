package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"simple-dapp/account"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial(os.Getenv("CHAIN_HOST"))
	if err != nil {
		panic(err)
	}

	privKey, err := crypto.HexToECDSA(os.Getenv("ACCOUNT_PRIV"))
	if err != nil {
		panic(err)
	}

	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		panic("Cannot cast pubkey to ECDSA")
	}

	addr := crypto.PubkeyToAddress(*pubKeyECDSA)

	c_addr := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
	acc, err := account.NewAccount(c_addr, client)
	if err != nil {
		panic(err)
	}

	count, err := acc.GetCount(nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(count)
	chain_id := new(big.Int)
	chain_id, ok = chain_id.SetString(os.Getenv("CHAIN_ID"), 10)
	if !ok {
		panic("Cannot convert chain ID")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chain_id)
	if err != nil {
		panic(err)
	}

	_, err = acc.IncrementCount(auth)
	if err != nil {
		panic(err)
	}

	count, err = acc.GetCount(nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(count)
}
