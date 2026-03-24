package main

import (
	"fmt"
	"log"

	"github.com/yukno1/ForeverStore/p2p"
)

func OnPeer(p p2p.Peer) error {
	fmt.Println("doing some logic with the peer outside of TCPTransport")
	p.Close()
	return nil
	// return fmt.Errorf("failed the onpeer func")
}

func main() {
	tcpOpts := p2p.TcpTransportOptions{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTcpTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
