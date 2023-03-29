package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	from	string
	to		string
	amount	int
}
type Block struct {
	Number		int
	time		time.Time
	txs			[]Transaction
	prevBlock	*Block
}
func (b *Block) AddTxs(txs... Transaction) {
	b.txs = append(b.txs, txs...)
}

func GetGenesisBlock() *Block {
	b := &Block {
		txs: nil,
		prevBlock: nil,
	}
	return b
}
func CreateBlock(number int, previousBlock *Block) *Block {
	b := GetGenesisBlock()
	b.Number = number
	b.time = time.Now()
	b.txs = []Transaction{}
	b.prevBlock = previousBlock
	return b
}

func main() {
	fmt.Println("OK")
	b1 := CreateBlock(1, nil)
	b2 := CreateBlock(2, b1)
	b1.AddTxs(Transaction{"1", "2", 11})
	fmt.Println(b1)
	fmt.Println(b1.time)
	fmt.Println(b2)
	fmt.Println(b2.time)
}
