package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("111111")
	bc.AddBlock("222222")
	bc.AddBlock("333333")

	it := bc.NewIterator()
	for {
		block := it.Next()

		fmt.Println("===========================\n\n")
		fmt.Printf("前去块哈希：%x\n", block.PreHash)
		fmt.Printf("当前区块哈希：%x\n", block.Hash)
		fmt.Printf("区快数据：%s\n", block.Data)

		if len(block.PreHash)==0{
			break
		}

	}

}
