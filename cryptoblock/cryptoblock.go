package cryptoblock

import (
	"bytes"
	"crypto/sha256"
)

type Cryptoblock struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (c *Cryptoblock) BuildHash() {
	details := bytes.Join([][]byte{c.Data, c.PrevHash}, []byte{})
	hash := sha256.Sum256(details)
	c.Hash = hash[:]
}

func BuildBlock(data string, prevHash []byte) *Cryptoblock {
	block := &Cryptoblock{[]byte{}, []byte(data), prevHash}
	block.BuildHash()
	return block
}
