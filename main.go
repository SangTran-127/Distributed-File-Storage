package main

import "DFS/p2p/transport"

func main() {
	tr := transport.NewTCP(":3000")

	if err := tr.Listen(); err != nil {

		panic(err)
	}

	select {}
}
