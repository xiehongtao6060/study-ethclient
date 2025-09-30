package main

import (
	"log"
	"study-ethclient/demo"

	//"study-ethclient/demo"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 请将此处替换为您的以太坊节点 URL (例如 Infura, Alchemy 或您自己的本地节点)
	//client, err := ethclient.Dial("wss://api.zan.top/node/ws/v1/eth/sepolia/0e10960fd7034a87807c5fec7b7531b3")
	client, err := ethclient.Dial("wss://api.zan.top/node/ws/v1/eth/mainnet/0e10960fd7034a87807c5fec7b7531b3")
	if err != nil {
		log.Fatal(err)
	}
	// 确保程序结束时关闭连接
	defer client.Close()
	// 运行示例 1: 查询区块
	demo.RunDemo1(client)
	// 运行示例 2: 查询交易详情
	//demo.RunDemo2(client)
}
