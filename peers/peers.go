package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func Unmarshal(peersBin []byte) ([]Peer, error) {
	if len(peersBin)%6 != 0 {
		err := fmt.Errorf("received malformed peers")
		return nil, err
	}

	var peers []Peer
	idx := 0
	for idx < len(peersBin) {
		var ip net.IP
		for i := 0; i < 4; i++ {
			ip = append(ip, peersBin[idx+i])
		}
		idx += 4

		var port uint16 = binary.BigEndian.Uint16(peersBin[idx : idx+2])
		idx += 2

		var p Peer = Peer{IP: ip, Port: port}
		peers = append(peers, p)
	}

	return peers, nil
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
