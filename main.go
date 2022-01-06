package main

import (
	// cr "crypto/rand"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	// ic "github.com/libp2p/go-libp2p-core/crypto"
	ws "github.com/libp2p/go-ws-transport"
)

func main() {

	peerkey, _, err := ic.GenerateEd25519Key(cr.Reader)

	opts := []libp2p.Option{
		//libp2p.ConnectionManager(connmgr.NewConnManager(2000, 3000, time.Minute)),
		libp2p.Identity(peerkey),
		//libp2p.BandwidthReporter(bwc),
		libp2p.Transport(ws.New),
		//libp2p.Transport(libp2pquic.NewTransport),
	}


	h, err := libp2p.New(opts...)
	if err != nil {
		panic(err)
	}

	for _, m := range h.Addrs() {
		fmt.Printf("%s/p2p/%s\n", m, h.ID())
	}


		ai, err := peer.AddrInfoFromString("/ip4/127.0.0.1/tcp/7878/ws/p2p/QmUez8XKFWJn8YQMjyoBjvURBWxZcw5LRiGWadiqDrNqd9")
		if err != nil {
			panic(err)
		}

		if err := h.Connect(context.TODO(), *ai); err != nil {
			panic(err)
		}

		s, err := h.NewStream(context.TODO(), ai.ID, "/bencher")
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(ioutil.Discard, s)
		if err != nil {
			panic(err)
		}



	return
}
