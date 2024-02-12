package blockchain

import (
	_block "github.com/AndroX7/go-eth/cryptoblock"
)

type BlockChain struct {
	blocks []*_block.Cryptoblock
}

func (c *BlockChain) AddBlock(data string) {
	prevBlock := c.blocks[len(c.blocks)-1]
	new := _block.BuildBlock(data, prevBlock.Hash)
	c.blocks = append(c.blocks, new)
}

func Inception() *_block.Cryptoblock {
	return _block.BuildBlock("Inception", []byte{})
}
