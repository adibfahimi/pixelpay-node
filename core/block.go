package core

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

const MAX_NONCE = ^uint64(0)

type Block struct {
	Difficulty   *big.Int `json:"difficulty"`
	Hash         string   `json:"hash"`
	PreviousHash string   `json:"previous_hash"`
	MerkleRoot   string   `json:"merkle_root"`
	Txs          []Tx     `json:"tx"`
	Index        uint     `json:"index"`
	Timestamp    uint     `json:"timestamp"`
	Nonce        uint64   `json:"nonce"`
}

func GenesisBlock() *Block {
	return &Block{
		Difficulty:   big.NewInt(4),
		Hash:         "",
		PreviousHash: "",
		MerkleRoot:   "",
		Txs:          []Tx{},
		Index:        0,
		Timestamp:    uint(time.Now().Unix()),
		Nonce:        0,
	}
}

func (b *Block) CalculateHash() string {
	b.MerkleRoot = b.CalculateMerkleRoot()

	return crypto.Keccak256Hash([]byte(fmt.Sprintf("%v%v%v%v%v", b.PreviousHash, b.MerkleRoot, b.Index, b.Timestamp, b.Nonce))).Hex()
}

func (b *Block) MineBlock() error {
	target := new(big.Int)
	target.Exp(big.NewInt(2), new(big.Int).Sub(big.NewInt(256), new(big.Int).SetUint64(b.Difficulty.Uint64())), nil)

	for {
		hash := b.CalculateHash()
		hashInt := new(big.Int)
		hashInt.SetString(hash[2:], 16)

		if hashInt.Cmp(target) == -1 {
			b.Hash = hash
			return nil
		}

		b.Nonce++

		if b.Nonce > MAX_NONCE {
			return errors.New("block mining failed")
		}
	}
}

func (b *Block) CalculateMerkleRoot() string {
	var hashes []string

	if len(b.Txs) == 0 {
		return ""
	}

	for _, tx := range b.Txs {
		hashes = append(hashes, tx.CalculateHash())
	}

	for len(hashes) > 1 {
		var levelHashes []string

		if len(hashes)%2 != 0 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}

		for i := 0; i < len(hashes); i += 2 {
			combined := fmt.Sprintf("%s%s", hashes[i], hashes[i+1])
			hash := fmt.Sprintf("%x", crypto.Keccak256([]byte(combined)))
			levelHashes = append(levelHashes, hash)
		}

		hashes = levelHashes
	}

	return hashes[0]
}

func (b *Block) IsValid() error {
	if b.Hash != b.CalculateHash() {
		return errors.New("invalid transaction hash")
	}

	if b.MerkleRoot != b.CalculateMerkleRoot() {
		return errors.New("invalid transaction merkle root")
	}

	if b.PreviousHash == "" {
		return errors.New("invalid transaction previous hash")
	}

	if b.Timestamp <= 0 {
		return errors.New("invalid transaction timestamp")
	}

	return nil
}

func (b *Block) Print() {
	fmt.Println(strings.Repeat("-", 33))

	fmt.Printf("Index: %v\n", b.Index)
	fmt.Printf("Hash: %v\n", b.Hash)
	fmt.Printf("Previous Hash: %v\n", b.PreviousHash)
	fmt.Printf("Index: %v\n", b.Index)
	fmt.Printf("Timestamp: %v\n", b.Timestamp)
	fmt.Printf("Nonce: %v\n", b.Nonce)

	for _, tx := range b.Txs {
		tx.Print()
	}

	fmt.Println(strings.Repeat("-", 33))
}
