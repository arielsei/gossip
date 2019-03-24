package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "Name", "Node name")
	uiPort := flag.Int("UIPort", 10000, "Node UI Port")
	gossipPort := flag.String("gossipPort", "127.0.0.1:5000", "NodeIP:NodeGossipPort")
	peers := flag.String("peers", "", "Peers known by node")
	flag.Parse()

	node, err := getNode(name, uiPort, gossipPort, peers)
	if err != nil {
		os.Exit(3)
	}
	fmt.Println(node.String())
}
