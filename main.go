package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PreHash []byte
	Hash []byte
	Data []byte
}
func NewBlock(preHash []byte,data string)*Block{
	block :=Block{
		PreHash:preHash,
		Hash:[]byte{},
		Data:[]byte(data),
	}
	block.SetHash()
	return &block
}
//生成哈希
func (block *Block)SetHash(){
	blockInfo:=append(block.PreHash,block.Data...)
	hash:=sha256.Sum256(blockInfo)
	block.Hash=hash[:]
}

func main(){
	block :=NewBlock([]byte{},"老师给班长红钻一个BTC")
	fmt.Printf("前去块哈希：%x\n",block.PreHash)
	fmt.Printf("当前区块哈希：%x\n",block.Hash)
	fmt.Printf("区块哈希值：%s\n",block.Data)
}
