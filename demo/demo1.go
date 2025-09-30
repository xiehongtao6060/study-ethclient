package demo

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func RunDemo1(client *ethclient.Client) {
	blockNumber := big.NewInt(5671744)
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功连接节点，获取到的最新区块高度为:", header.Number.String())

	// 使用获取到的最新区块号来查询完整的区块信息
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		log.Fatal(err)
	}

	// --- 使用更详细的描述来打印区块信息 ---
	fmt.Println("\n--- 开始打印完整区块信息 ---")
	fmt.Println("区块号 (Block Number):", block.Number().Uint64())
	fmt.Println("区块哈希 (Block Hash):", block.Hash().Hex())
	fmt.Println("区块时间戳 (Timestamp):", block.Time())
	fmt.Println("挖出此块的矿工地址 (Miner):", block.Coinbase().Hex())
	fmt.Println("区块中的交易数量 (Transaction Count):", len(block.Transactions()))
	fmt.Println("--- 完整区块信息打印完毕 ---\n")
}
