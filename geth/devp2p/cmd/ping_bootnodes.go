package main

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
)

func startTestServer(name string, port int, nodisc, nodial bool) *p2p.Server {
	config := p2p.Config{
		Name:        name,
		MaxPeers:    1,
		ListenAddr:  fmt.Sprintf("127.0.0.1:%d", port),
		NoDiscovery: nodisc,
		NoDial:      nodial,
		PrivateKey:  newkey(),
	}

	server := &p2p.Server{
		Config: config,
	}

	if err := server.Start(); err != nil {
		panic("Could not start server: " + err.Error())
	}

	return server
}

func newkey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic("couldn't generate key: " + err.Error())
	}

	return key
}

func syncAddPeer(srv *p2p.Server, node *enode.Node) bool {
	var (
		ch      = make(chan *p2p.PeerEvent)
		sub     = srv.SubscribeEvents(ch)
		timeout = time.After(10 * time.Second)
	)
	defer sub.Unsubscribe()
	srv.AddPeer(node)
	for {
		select {
		case ev := <-ch:
			if ev.Type == p2p.PeerEventTypeAdd && ev.Peer == node.ID() {
				return true
			}
		case <-timeout:
			return false
		}
	}
}

func newNode(url string) (*enode.Node, error) {
	node, err := enode.Parse(enode.ValidSchemes, url)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func pingBootnodes(srv *p2p.Server, bootnodes []string) []string {
	connected := make([]string, 0)
	for _, v := range bootnodes {
		node, err := newNode(v)
		if err != nil {
			fmt.Println("failed to parse", v)
		}
		if syncAddPeer(srv, node) {
			connected = append(connected, v)
		}
	}
	return connected
}

func main() {
	srv1 := startTestServer("server-1", 30301, false, false)
	defer srv1.Stop()

	connected := pingBootnodes(srv1, params.MainnetBootnodes)
	fmt.Printf("Connected \033[35m%d\033[0m Mainnet Bootnode(s): \033[36m%v\033[0m\n", len(connected), connected)

	connected = pingBootnodes(srv1, params.HoodiBootnodes)
	fmt.Printf("Connected \033[35m%d\033[0m Hoodi Bootnode(s): \033[36m%v\033[0m\n", len(connected), connected)

	connected = pingBootnodes(srv1, params.HoleskyBootnodes)
	fmt.Printf("Connected \033[35m%d\033[0m Holesky Bootnode(s): \033[36m%v\033[0m\n", len(connected), connected)

	connected = pingBootnodes(srv1, params.SepoliaBootnodes)
	fmt.Printf("Connected \033[35m%d\033[0m Sepolia Bootnode(s): \033[36m%v\033[0m\n", len(connected), connected)

	connected = pingBootnodes(srv1, params.V5Bootnodes)
	fmt.Printf("Connected \033[35m%d\033[0m V5 Bootnode(s): \033[36m%v\033[0m\n", len(connected), connected)
}
