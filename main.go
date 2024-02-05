package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/AndroX7/go-eth/config"
	"github.com/AndroX7/go-eth/driver"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	cfg, _ := config.Get("./.env")
	err := driver.Initialize(cfg)

	if err != nil {
		panic(err)
	}

	blockNumber := big.NewInt(5532993)
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := driver.Instance().Client().BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}
