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
//3.生成哈希
func (block *Block)SetHash(){
	blockInfo:=append(block.PreHash,block.Data...)
	hash:=sha256.Sum256(blockInfo)
	block.Hash=hash[:]
}

//4.引入区块连
type BlockChain struct {
	blocks []*Block
}
//5.定义一个区块连
func NewBlockChain()*BlockChain  {
	genesisBlock:=GenesisBlock()
	 return &BlockChain{
	 	blocks:[]*Block{genesisBlock},
	 }
}
//6.定义创世区块
func GenesisBlock()*Block{
	return NewBlock([]byte{},"创世区块，真牛鼻！")
}

func main(){
     bc:=NewBlockChain()
     for i,block:=range bc.blocks{
     	 fmt.Printf("=======当前区块高度:%d======\n",i)
		 fmt.Printf("前去块哈希：%x\n",block.PreHash)
		 fmt.Printf("当前区块哈希：%x\n",block.Hash)
		 fmt.Printf("区块哈希值：%s\n",block.Data)
	 }


}
