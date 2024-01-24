package core

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Tx struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Signature string `json:"signature"`
	Hash      string `json:"hash"`
	Amount    uint   `json:"amount"`
	Timestamp uint   `json:"timestamp"`
}

func NewTx(from string, to string, amount uint, signature string) *Tx {
	tx := Tx{
		From:      from,
		To:        to,
		Signature: signature,
		Hash:      "",
		Amount:    amount,
		Timestamp: uint(time.Now().Unix()),
	}

	tx.Hash = tx.CalculateHash()

	return &tx
}

func (tx *Tx) CalculateHash() string {
	return crypto.Keccak256Hash([]byte(fmt.Sprintf("%s%s%d%d", tx.From, tx.To, tx.Amount, tx.Timestamp))).Hex()
}

func (tx *Tx) IsValid() error {
	if tx.From == "" {
		return errors.New("invalid block from is empty")
	}

	if tx.To == "" {
		return errors.New("invalid block to is empty")
	}

	if tx.Amount <= 0 {
		return errors.New("invalid block amount is invalid")
	}

	if tx.Signature == "" {
		return errors.New("invalid block signature is empty")
	}

	if tx.Hash == "" {
		return errors.New("invalid block hash is empty")
	}

	if tx.Timestamp <= 0 {
		return errors.New("invalid block timestamp is invalid")
	}

	if tx.Hash != tx.CalculateHash() {
		return errors.New("invalid block hash is invalid")
	}

	if err := tx.VerifySignature(); err != nil {
		return err
	}

	return nil
}

func (tx *Tx) VerifySignature() error {
	pubKeyBytes, err := hexutil.Decode(tx.From)
	if err != nil {
		return fmt.Errorf("failed to decode: %v", err)
	}

	signatureBytes, err := hexutil.Decode(tx.Signature)
	if err != nil {
		return fmt.Errorf("failed to decode: %v", err)
	}

	hashBytes, err := hexutil.Decode(tx.Hash)
	if err != nil {
		return fmt.Errorf("failed to decode: %v", err)
	}

	realPubKey, err := crypto.SigToPub(hashBytes, signatureBytes)
	if err != nil {
		return fmt.Errorf("failed to recover public key: %v", err)
	}

	realPubKeyBytes := crypto.CompressPubkey(realPubKey)

	if bytes.Equal(pubKeyBytes, realPubKeyBytes) {
		return fmt.Errorf("invalid public key")
	}

	if !crypto.VerifySignature(realPubKeyBytes, hashBytes, signatureBytes[:64]) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

func (tx *Tx) Print() {
	fmt.Println(strings.Repeat("-", 33))

	fmt.Printf("Hash: %v\n", tx.Hash)
	fmt.Printf("From: %v\n", tx.From)
	fmt.Printf("To: %v\n", tx.To)
	fmt.Printf("Signature: %v\n", tx.Signature)
	fmt.Printf("Amount: %v\n", tx.Amount)
	fmt.Printf("Timestamp: %v\n", tx.Timestamp)

	fmt.Println(strings.Repeat("-", 33))
}
