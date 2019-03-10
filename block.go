package main

import (
	"crypto/sha256"
	"time"
)

type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PreHash []byte
	//3.美克尔根 merkeltree
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数
	Nonce uint64

	//a.当前区块哈希
	Hash []byte
	//b.数据
	Data []byte
}
func NewBlock(preHash []byte,data string)*Block{
	block :=Block{
		Version:00,
		PreHash:preHash,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0,
		Nonce:0,
		Hash:[]byte{},
		Data:[]byte(data),
	}
	block.SetHash()
	return &block
}
//辅助函数，将uint64转换成[]byte
func Uint64ToBtye(num uint64)[]byte{
	return []byte{}
}
//3.生成哈希
func (block *Block)SetHash(){
	var blockInfo []byte

	blockInfo=append(blockInfo,Uint64ToBtye(block.Version)...)
	blockInfo=append(blockInfo,block.PreHash...)
	blockInfo=append(blockInfo,block.MerkelRoot...)
	blockInfo=append(blockInfo,Uint64ToBtye(block.TimeStamp)...)
	blockInfo=append(blockInfo,Uint64ToBtye(block.Difficulty)...)
	blockInfo=append(blockInfo,Uint64ToBtye(block.Nonce)...)
	blockInfo=append(blockInfo,block.Data...)

	hash:=sha256.Sum256(blockInfo)
	block.Hash=hash[:]
}
