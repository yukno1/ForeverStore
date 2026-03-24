package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {

	opts := TcpTransportOptions{
		ListenAddr:    ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTcpTransport(opts)

	assert.Equal(t, tr.ListenAddr, opts.ListenAddr)

	// server
	//tr.Start()

	assert.Nil(t, tr.ListenAndAccept())

}
