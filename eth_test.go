package geth

import (
	"fmt"
	"testing"
)

func TestEthClient_EthAccount(t *testing.T) {
	cli = NewEthClient(mainNetwork)

	fmt.Println("eth client dial success...😊😊😊")
	fmt.Println(cli.r)

	cli.EthAccount()
}
