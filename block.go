package main

import (
	"crypto/sha256"
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
