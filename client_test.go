package geth

import (
	"fmt"
	"testing"
)

var (
	mainNetwork = "https://cloudflare-eth.com"
)

var cli *EthClient

func TestNewClientEth(t *testing.T) {

	cli = NewEthClient(mainNetwork)

	fmt.Println("eth client dial success...😊😊😊")
	fmt.Println(cli.r)
	
}
