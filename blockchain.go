package main

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
//7.添加区块实现
func (bc *BlockChain)AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	block := NewBlock(prevHash, data)
	bc.blocks = append(bc.blocks, block)
}