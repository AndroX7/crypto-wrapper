package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/AndroX7/crypto-wrapper/config"
	"github.com/AndroX7/crypto-wrapper/driver"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	cfg, _ := config.Get("./.env")
	err := driver.Initialize(cfg)

	if err != nil {
		panic(err)
	}

	blockNumber := big.NewInt(5532993)
	client := driver.Instance().Client()
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
	if client != nil {
		client.Close()
	}
	fmt.Println("done")
}
