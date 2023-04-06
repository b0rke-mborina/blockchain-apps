package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
)

const AddressLength = 4
type Hash [32]byte
type Address [AddressLength]byte

type Transaction struct {
	From		Address
	To			Address
	Amount		int
	Signature	[]byte
}
func NewTransaction(from Address, to Address, amount int) *Transaction {
	b :=  &Transaction {
		From:	from,
		To:		to,
		Amount:	amount,
		Signature: nil,
	}
	return b
}
func GenerateNewAddress() (a Address) {
	addr := make([]byte, 4)
	rand.Read(addr)
	copy(a[:], addr)
	return a
}

func (t *Transaction) TxHash() Hash  {
	return ValuesHash(
		t.From[:],
		t.To[:],
		[]byte( fmt.Sprintf("%d", t.Amount)[:]),
	)
}
func ValuesHash(a ...[]byte) (retval Hash)  {
	buffer := new(bytes.Buffer)
	for _, val := range a {
		buffer.Write(val)
	}
	retval = sha256.Sum256(buffer.Bytes())
	return
}

func (t *Transaction) Sign(key ecdsa.PrivateKey) {
	txHash := t.TxHash()
	r, s, err := ecdsa.Sign(rand.Reader, &key, txHash[:])
	if err != nil {
		panic(err)
	}
	t.Signature = append(r.Bytes()[:32], s.Bytes()[:32]...)
}

func (t* Transaction) Verify(key ecdsa.PublicKey) bool {
	txHash := t.TxHash()
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(t.Signature[:32])
	s.SetBytes(t.Signature[32:64])
	isValid := ecdsa.Verify(&key, txHash[:], r, s)
	return isValid
}

func main() {
	tx := NewTransaction(GenerateNewAddress(), GenerateNewAddress(), 20)

	pubKeyCurve := elliptic.P256() // http://golang.org/pkg/crypto/elliptic/#P256

	privateKey := new(ecdsa.PrivateKey)
	privateKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pubKey := privateKey.PublicKey

	// potpiši transakciju
	tx.Sign(*privateKey)

	// verificiraj
	fmt.Println("Valid? ", tx.Verify(pubKey))
}