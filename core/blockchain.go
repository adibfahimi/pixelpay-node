package core

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

var BC = NewBlockchain()

type Blockchain struct {
	Difficulty   *big.Int `json:"difficulty"`
	Chain        []Block  `json:"chain"`
	PendingBlock Block    `json:"pending_block"`
	MiningReward uint     `json:"mining_reward"`
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenesisBlock()

	return &Blockchain{
		Chain: []Block{*genesisBlock},
		PendingBlock: Block{
			Hash:         "",
			PreviousHash: genesisBlock.Hash,
			MerkleRoot:   "",
			Txs:          []Tx{},
			Index:        1,
			Timestamp:    uint(time.Now().Unix()),
			Nonce:        0,
			Difficulty:   big.NewInt(4),
		},
		Difficulty:   big.NewInt(4),
		MiningReward: 100,
	}
}

func (bc *Blockchain) GetLatestBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) Print() {
	fmt.Println(strings.Repeat("-", 33))

	fmt.Printf("Chain: %v\n", bc.Chain)
	bc.PendingBlock.Print()
	fmt.Printf("Difficulty: %v\n", bc.Difficulty)
	fmt.Printf("Mining Reward: %v\n", bc.MiningReward)
	for _, block := range bc.Chain {
		block.Print()
	}

	fmt.Println(strings.Repeat("-", 33))
}
