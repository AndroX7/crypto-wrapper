package blockdata

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type block struct {
	Data         map[string]interface{}
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Pow          int
}

type Block interface {
	CalculateHash() string
	Mine(difficulty int)
	GetBlock() *block
}

func New(data interface{}, hash, prevHash string, timestamp time.Time, pow int) Block {
	return &block{
		Data:         data.(map[string]interface{}),
		Hash:         hash,
		PreviousHash: prevHash,
		Timestamp:    timestamp,
		Pow:          pow,
	}
}

func (b *block) CalculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.Pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *block) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = b.CalculateHash()
	}
}

func (b *block) GetBlock() *block {
	return b
}
