package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {

	opts := TcpTransportOptions{
		ListenAddr: ":4000",
	}
	tr := NewTcpTransport(opts)

	assert.Equal(t, tr.ListenAddr, opts.ListenAddr)

	// server
	//tr.Start()

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
