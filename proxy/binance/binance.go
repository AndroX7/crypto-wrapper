package binance

import (
	client "github.com/adshao/go-binance/v2"
	delivery "github.com/adshao/go-binance/v2/delivery"
	future "github.com/adshao/go-binance/v2/futures"
)

// var client.n
type Binance struct {
	Client *client.Client
	USDT   *future.Client
	COINM  *delivery.Client
}

func New(apiKey, secretKey string) *Binance {
	return &Binance{
		Client: client.NewClient(apiKey, secretKey),
		USDT:   future.NewClient(apiKey, secretKey),
		COINM:  delivery.NewClient(apiKey, secretKey),
	}
}
