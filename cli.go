package main

import (
	"fmt"
	"os"
)

//这是一个接受命令行参数并且控制区块连操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
   addBlock --data DATA    "add data to blockchain"
   printChain              "print all blockchain data"

`

func (c *CLI) Run() {
	args := os.Args

	if len(args) < 2 {
		println(Usage)
		return
	}

	cmd := args[1]
	switch cmd {

	case "addBlock":
		//添加区块
		fmt.Println("添加区块")
	case "printChain":
		//打印区块
		fmt.Println("打印区块")
	default:
		fmt.Println(Usage)
		return
	}
}
