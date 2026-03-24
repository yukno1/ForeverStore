package p2p

import "net"

// Message holds any arbitrary data being sent over
// each transport btw 2 nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
