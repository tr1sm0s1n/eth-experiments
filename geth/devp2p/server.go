package main

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
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
		timeout = time.After(2 * time.Second)
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

func main() {
	srv1 := startTestServer("server-1", 30301, false, false)
	defer srv1.Stop()
	fmt.Printf("Server 1 Enode: \033[36m%s\033[0m\n", srv1.NodeInfo().Enode)

	srv2 := startTestServer("server-2", 30302, false, true)
	defer srv2.Stop()
	fmt.Printf("Server 2 Enode: \033[36m%s\033[0m\n", srv2.NodeInfo().Enode)

	if !syncAddPeer(srv1, srv2.Self()) {
		panic("peer not connected")
	}
	fmt.Printf("Server 1 Peer Count: \033[36m%d\033[0m\n", srv1.PeerCount())

	fmt.Printf("\033[31mRemoving peer..\033[0m\n")
	srv1.RemovePeer(srv2.Self())

	fmt.Printf("Server 1 Peer Count: \033[36m%d\033[0m\n", srv1.PeerCount())
}
