package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Data		string
	prevNode	string
}
func (n *Node) addElement(data string) *Node {
	node :=  &Node {
		Data:	data,
		prevNode: n.Hash(),
	}
	return node
}
func (n *Node) Hash() string {
	buffer := new(bytes.Buffer)
	buffer.Write([]byte(n.Data))
	buffer.Write([]byte(n.prevNode))
	result := sha256.Sum256(buffer.Bytes())
	return string(result[:])
}


func main() {
	a := Node{Data: "<GENESIS>"}
	b := a.addElement("a")
	c := b.addElement("b")

	a1 := Node{Data: "<GENESIS>"}
	b1 := a1.addElement("a")
	c1 := b1.addElement("b")

	// hash mora biti jednak za oba lanca
	fmt.Printf("Prvi blockchain: %s \n", c.Hash())
	fmt.Printf("Drugi blockchain: %s \n", c1.Hash())
	// fmt.Println()
	// fmt.Println(c.Hash())
	// fmt.Println(c1.Hash())
	// fmt.Println()
	// fmt.Println(strings.Compare(c.Hash(), c1.Hash()))
}