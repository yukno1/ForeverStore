package p2p

// ErrInvalidHandShake is returned if the handshake btw
// the local and remote node could not be established
// var ErrInvalidHandShake = errors.New(("invalid handshake"))

type Handshaker interface {
	Handshake() error
}

// HandshakeFunc is 
type HandshakeFunc func(Peer) error 

type DefaultHandshaker struct {

}

func NOPHandshakeFunc(Peer) error {
	return nil
}