package main

import (
	"bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	currentHashPointer []byte

}

func (bc *BlockChain)NewIterator()*BlockChainIterator{
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

func (it *BlockChainIterator)Next()*Block{
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
        bucket:=tx.Bucket([]byte(blockBucket))
        if bucket == nil{
        	log.Panic("迭代器错误")
		}
        tempblock := bucket.Get(it.currentHashPointer)
        block = Deserialize(tempblock)

		it.currentHashPointer=block.PreHash

		return nil
	})

	return &block

}