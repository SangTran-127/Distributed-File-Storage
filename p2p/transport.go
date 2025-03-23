package p2p

type Peer interface {
}

// Transport Can be TCP, UDP, Websocket
type Transport interface {
	Listen(address string)
}
