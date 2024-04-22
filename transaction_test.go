package geth

import (
	"fmt"
	"testing"
)

func TestEthClient_EthBlockTransaction(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	var blockNumber int64 = 5671744
	transactions := cli.EthBlockTransaction(blockNumber)

	for ids, transaction := range transactions {
		fmt.Println(transaction.Hash) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(ids)              // 0
	}
}

func TestEthClient_EthBlockSingleTransactionByHash(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	var blockNumber int64 = 5671744
	transactions := cli.EthBlockTransaction(blockNumber)

	for _, transaction := range transactions {
		tx, isPending := cli.EthBlockSingleTransactionByHash(transaction.Hash.String())
		fmt.Println(isPending)
		fmt.Println(tx)
	}

}
