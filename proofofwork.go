package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义一个工作量证明
type ProofOfWork struct {
	block *Block
	target *big.Int
}

//提供创建POW的函数
func NewProofOfWork(block *Block)*ProofOfWork{
	pow :=ProofOfWork{
		block:block,
	}
	targetStr :="0000100000000000000000000000000000000000000000000000000000000000"
	tempInt :=big.Int{}
	tempInt.SetString(targetStr,16)

	pow.target=&tempInt
	return &pow
}

func (pow *ProofOfWork)Run()([]byte,uint64){
    var nonce uint64
    block :=pow.block
    var hash [32]byte

	for   {
		temp:=[][]byte{
			Uint64ToBtye(block.Version),
			block.PreHash,
			block.MerkelRoot,
			Uint64ToBtye(block.TimeStamp),
			Uint64ToBtye(block.Difficulty),
			Uint64ToBtye(nonce),
			block.Data,
		}
		blockInfo:=bytes.Join(temp,[]byte{})
		hash = sha256.Sum256(blockInfo)

		tempInt:=big.Int{}
		tempInt.SetBytes(hash[:])

		if tempInt.Cmp(pow.target)==-1 {
			fmt.Printf("挖矿成功！hash :%x,nonce:%d\n",hash,nonce)
			return hash[:],nonce
		}else {
			nonce ++
		}
	}
}
