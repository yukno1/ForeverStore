package p2p

// Message holds any arbitrary data being sent over
// each trransport btw 2 nodes
type Message struct {
	Payload []byte
}