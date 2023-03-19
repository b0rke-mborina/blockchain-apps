package main

import (
	"fmt"
	"time"
)

// 1.
type item struct {
	id 		int
	name	string
	price	int
}
type game struct {
	item
	genre	string
}

// 2.
type Transaction struct {
	from	string
	to		string
	amount	int
}
type Block struct {
	Number			int
	time			time.Time
	transactions	[]Transaction
}
func newBlock(number int) *Block {
	b :=  &Block {
		Number:			number,
		time:			time.Now(),
		transactions:	[]Transaction{},
	}
	return b
}
func (b *Block) addTransaction(txs... Transaction) {
	b.transactions = append(b.transactions, txs...)
}
func (b *Block) String() string {
	return fmt.Sprint("Block\nNumber: ", b.Number, "\nTime: ", b.time, "\nTransactions: ", b.transactions)
}

func main() {
	fmt.Println("OK")

	// 1.
	data := []game{
		game{item{1, "god of war", 50}, "action adventure"},
		game{item{2, "x-com 2", 30}, "strategy"},
		game{item{3, "minecraft", 20}, "sandbox"},
	}
	for index, game := range data {
		fmt.Println("GAME", index + 1)
		fmt.Println("Item ID:", game.id)
		fmt.Println("Item name:", game.name)
		fmt.Println("Item price:", game.price)
		fmt.Println("Game genre:", game.genre)
		fmt.Println()
	}

	// 2.
	block := newBlock(4)
	block.addTransaction(Transaction{"1", "2", 11})
	block.addTransaction(Transaction{"3", "4", 55}, Transaction{"4", "5", 66})
	fmt.Println(block)
}
