package main

import "github.com/ethereum/go-ethereum/p2p/enode"

func NewNode(url string) (*enode.Node, error) {
	node, err := enode.Parse(enode.ValidSchemes, url)
	if err != nil {
		return nil, err
	}
	return node, nil
}
