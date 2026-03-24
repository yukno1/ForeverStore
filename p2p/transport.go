package p2p

// interface represent remote node
type Peer interface {
	Close() error
}

// handle comm btw nodes in networks
// (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
