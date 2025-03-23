package transport

import (
	"DFS/p2p"
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	connection net.Conn
	// Client Dial to make connection -> outbound(send)
	outbound bool
	// Server Listen and Accept connection -> inbound(received)
	inbound bool
}

type TCP struct {
	address          string
	listener         net.Listener
	handShakeHandler HandShakeHandler
	decoder          p2p.Decoder
	// Protected TCP prevent conflict
	mu    sync.Mutex
	peers map[string]p2p.Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		connection: conn,
		outbound:   outbound,
	}
}

func NewTCP(address string) *TCP {
	return &TCP{
		address:          address,
		handShakeHandler: NOPHandShakeHandler,
	}
}

func (t *TCP) Listen() error {
	var err error
	t.listener, err = net.Listen("tcp", t.address)
	if err != nil {
		return err
	}

	go t.acceptListener()
	return nil
}

func (t *TCP) acceptListener() {
	for {
		connection, err := t.listener.Accept()
		if err != nil {
			fmt.Printf(err.Error())
		}

		go t.connectionHandler(connection)
	}
}

func (t *TCP) connectionHandler(connection net.Conn) {
	peer := NewTCPPeer(connection, true)

	if err := t.handShakeHandler(peer.connection); err != nil {
		fmt.Printf(err.Error())
		connection.Close()
	}

	fmt.Printf("New connection from %s\n", peer)
}
