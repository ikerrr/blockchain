package main

import (
	"bolt"
	"fmt"
	"log"
)

//4.引入区块连
type BlockChain struct {
	//blocks []*Block
	db *bolt.DB
	tail []byte
}

const blockChainDb  = "blockChain.db"
const blockBucket  ="blockBucket"



//5.定义一个区块连
func NewBlockChain()*BlockChain  {
	//return &BlockChain{
	//	blocks:[]*Block{genesisBlock},
	//}
	var lastHash []byte

	db,err:=bolt.Open(blockChainDb,0600,nil)
	defer db.Close()
	if err!=nil{
		log.Panic("打开数据库失败")
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket==nil{

			bucket,err = tx.CreateBucket([]byte(blockBucket))
			if err !=nil{
				fmt.Println("bucket创建失败")
			}

			genesisBlock:=GenesisBlock()


			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"),genesisBlock.Hash)
			lastHash=genesisBlock.Hash
		}else {
		    lastHash = bucket.Get([]byte("LastHash"))
		}
		return nil

	})

     return &BlockChain{db,lastHash}
}



//6.定义创世区块
func GenesisBlock()*Block{
	return NewBlock([]byte{},"创世区块，真牛鼻！")
}
//7.添加区块实现
func (bc *BlockChain)AddBlock(data string) {
/*	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	block := NewBlock(prevHash, data)
	bc.blocks = append(bc.blocks, block)
*/
}