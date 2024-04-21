package geth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	EthClientDialError = "Dial eth client error"
)

type EthClient struct {
	r *ethclient.Client
}

func NewEthClient(network string) *EthClient {

	c, err := ethclient.Dial(network)
	if err != nil {
		log.Fatal(EthClientDialError + "【😭】" + err.Error())
	}
	e := &EthClient{
		r: c,
	}

	return e
}
