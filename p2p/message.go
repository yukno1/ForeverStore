package p2p

import "net"

// RPC holds any arbitrary data being sent over
// each transport btw 2 nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}
