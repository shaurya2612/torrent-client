package client

import (
	"net"
	"time"

	"github.com/shaurya2612/torrent-client/peers"
	"github.com/veggiedefender/torrent-client/bitfield"
)

// A Client is a TCP connection with a peer
type Client struct {
	Conn     net.Conn
	Choked   bool
	Bitfield bitfield.Bitfield
	peer     peers.Peer
	infoHash [20]byte
	peerID   [20]byte
}

func New(peer peers.Peer, peerID, infoHash [20]byte) (*Client, error) {
	conn, err := net.DialTimeout("tcp", peer.String(), 3*time.Second)

	if err != nil {
		return nil, err
	}
}
