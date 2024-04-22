package geth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

var (
	InquireEthBlockNumberError               = "Inquire eth block number failed !!!"
	InquireEthBlockByNumberError             = "Inquire eth block by number failed !!!"
	InquireEthBlockTransactionCountHashError = "Inquire eth block transaction count by hash failed !!!"
)

// EthBlockNumber 最新区块高度
func (e *EthClient) EthBlockNumber() string {
	number, err := e.r.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("【%s】%s", InquireEthBlockNumberError, err)
	}
	return number.Number.String()
}

// EthBlockNumber2 最新区块高度
func (e *EthClient) EthBlockNumber2() uint64 {
	number, err := e.r.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("【%s】%s", InquireEthBlockNumberError, err)
	}
	return number
}

// EthBlockContentByBlockNumber 读取区块的所有内容和元数据
func (e *EthClient) EthBlockContentByBlockNumber(blockNumber int64) *types.Block {
	n := big.NewInt(blockNumber)
	block, err := e.r.BlockByNumber(context.Background(), n)
	if err != nil {
		log.Fatalf("【%s】%s", InquireEthBlockByNumberError, err)
	}
	return block
}

// EthBlockTransactionCountByBlockNumber 块中的交易计数
func (e *EthClient) EthBlockTransactionCountByBlockNumber(hash common.Hash) uint {
	count, err := e.r.TransactionCount(context.Background(), hash)
	if err != nil {
		log.Fatalf("【%s】%s", InquireEthBlockTransactionCountHashError, err)
	}
	return count
}
