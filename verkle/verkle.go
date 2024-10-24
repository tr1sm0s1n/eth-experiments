package main

import (
	"bytes"
	"encoding/hex"
	"log"

	"github.com/ethereum/go-verkle"
)

var (
	// a 32 byte value, as expected in the tree structure
	testValue      = []byte("0123456789abcdef0123456789abcdef")
	zeroKeyTest, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
)

func main() {
	v := verkle.New()

	if err := v.Insert(zeroKeyTest, testValue, nil); err != nil {
		log.Fatalf("error inserting: %v", err)
	}

	leaf, ok := v.(*verkle.InternalNode).Children()[0].(*verkle.LeafNode)
	if !ok {
		log.Fatalf("invalid leaf node type")
	}

	if !bytes.Equal(leaf.Values()[zeroKeyTest[verkle.StemSize]], testValue) {
		log.Fatalf("did not find correct value in trie")
	}

	log.Println("Verkle tree leaf verified!!")
}
