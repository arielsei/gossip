package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Node of Gossip system
type Node struct {
	UIPort     int
	ip         string
	gossipPort int
	name       string
	peers      []*Node
}

func (node *Node) String() string {
	return fmt.Sprintf("Node %s:\nUI Port: %s\nGossip Port: %d\nUI Port: %d\nPeers: [%s]",
		node.name, node.ip, node.gossipPort, node.UIPort, node.PeersString())
}

func (node *Node) PeersString() string {
	var peersStr string
	for _, node := range node.peers {
		peersStr += node.GossipPortString() + ", "
	}
	return peersStr
}

func (node *Node) GossipPortString() string {
	return fmt.Sprintf("%s:%d", node.ip, node.gossipPort)
}

func getNode(name *string, uiPort *int, gossipPort,
	peersStr *string) (*Node, error) {
	var node Node

	ip, port, err := getIPPort(*gossipPort)
	if err != nil {
		return nil, err
	}

	node.ip = ip
	node.gossipPort = port

	if name == nil {
		node.name = ""
	} else {
		node.name = *name
	}

	if uiPort == nil {
		node.UIPort = 5001
	} else {
		node.UIPort = *uiPort
	}

	if peersStr != nil {
		peers, err := getPeers(*peersStr)
		if err != nil {
			return nil, err
		}
		node.peers = peers
	}

	return node.Validate()
}

func (node *Node) Validate() (*Node, error) {
	if node.gossipPort <= 5000 || node.UIPort <= 5000 {
		return nil, errors.New("Node ports should be higher than 5000")
	}

	reIP := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if !reIP.MatchString(node.ip) {
		return nil, errors.New("Node ip is not well formed")
	}

	return node, nil
}

func getIPPort(s string) (string, int, error) {
	info := strings.Split(s, ":")
	port, err := strconv.Atoi(Last(info))

	return info[0], port, err
}

func getPeers(s string) ([]*Node, error) {
	nodeStrs := strings.Split(s, "_")
	var nodes []*Node
	for _, nodeStr := range nodeStrs {
		node, err := getNode(nil, nil, &nodeStr, nil)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
