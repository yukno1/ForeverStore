package p2p

// interface represent remote node
type Peer interface {
}

// handle comm btw nodes in networks
// (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
