package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

func makePeer(addr string) (*peerstore.PeerInfo, error) {
	maddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	return peerstore.InfoFromP2pAddr(maddr)
}

func dialRelayNode(ctx context.Context, h host.Host, relayAddr string) error {
	p, err := makePeer(relayAddr)
	if err != nil {
		return err
	}

	return h.Connect(ctx, *p)
}
