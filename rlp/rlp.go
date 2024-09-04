package main

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	data := "Hello, Ethereum!"
	fmt.Println("Original data:", data)

	encoded, err := rlp.EncodeToBytes(data)
	if err != nil {
		panic(fmt.Sprintf("err: %s", err))
	}

	encodedHex := hex.EncodeToString(encoded)
	fmt.Println("Encoded data:", encodedHex)

	ptr := new(string)

	if err = rlp.DecodeBytes(encoded, ptr); err != nil {
		panic(fmt.Sprintf("err: %s", err))
	}

	decoded := reflect.ValueOf(ptr).Elem().String()
	if decoded != data {
		fmt.Println("Oops, data mismatch.")
	}

	fmt.Println("Data matched!!")
	fmt.Println("Decoded data:", decoded)
}
