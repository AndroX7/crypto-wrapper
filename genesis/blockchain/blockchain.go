package blockchain

import (
	"time"

	block "github.com/AndroX7/crypto-wrapper/genesis/blockdata"
)

type blockchain struct {
	genesisBlock *block.Block
	chain        []block.Block
	minners      int
}

type Blockchain interface {
	AddBlock(from, to string, amount float64)
	IsValid() bool
}

var genesis *block.Block

func New(minners int) Blockchain {
	if genesis == nil {
		tmp := block.New(nil, "0", "", time.Now().UTC(), minners)
		genesis = &tmp
	}
	return &blockchain{
		genesis,
		[]block.Block{*genesis},
		minners,
	}
}

func (b *blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	tmp := lastBlock.GetBlock()
	new := block.New(blockData, "", tmp.PreviousHash, time.Now().UTC(), 0)
	newBlock := new.GetBlock()
	newBlock.Mine(b.minners)
	b.chain = append(b.chain, newBlock)
}

func (b *blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		block := b.chain[i]
		prevChain := block.GetBlock()
		block = b.chain[i+1]
		currentChain := block.GetBlock()
		if currentChain.Hash != currentChain.CalculateHash() || currentChain.PreviousHash != prevChain.Hash {
			return false
		}
	}
	return true
}
