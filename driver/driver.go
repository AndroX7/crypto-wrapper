package driver

import (
	"os"

	"github.com/AndroX7/crypto-wrapper/config"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Driver struct {
	client      *ethClient.Client
	logger      log.Logger
	initialized bool
}

var instance *Driver

func Instance() *Driver {
	if instance == nil {
		instance = &Driver{}
	}
	return instance
}

func Initialize(config config.Configuration) error {
	var tmp log.Logger
	if config.LogOutput == "stderr" {
		tmp = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	} else {
		tmp = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	}
	tmp = log.With(tmp, "ts", log.DefaultTimestampUTC)
	if config.LogType == "debug" {
		tmp = level.NewFilter(tmp, level.AllowDebug())
	} else if config.LogType == "info" {
		tmp = level.NewFilter(tmp, level.AllowInfo())
	} else if config.LogType == "error" {
		tmp = level.NewFilter(tmp, level.AllowError())
	} else if config.LogType == "none" {
		tmp = level.NewFilter(tmp, level.AllowNone())
	}
	Instance().logger = tmp

	client, err := ethClient.Dial(config.RawURL)
	if err != nil {
		return err
	}
	Instance().client = client
	Instance().initialized = true
	return nil
}

func (d *Driver) Client() *ethClient.Client {
	return d.client
}
