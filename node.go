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
	uiPort     int
	ip         string
	gossipPort int
	name       string
	peers      []*Node
}

func (node *Node) String() string {
	return fmt.Sprintf("Node %s:\nUIPort: %s\n", node.name, node.ip)
}

func (node *Node) UIIPPort() string {
	if len(node.ip) == 0 || node.uiPort <= 5000 {
		return ""
	}

	return fmt.Sprintf("%s:%d", node.ip, node.uiPort)
}

func getNode(name *string, uiPort *int, gossipPort,
	peers *string) (*Node, error) {
	var node Node

	ip, port, err := getIPPort(*gossipPort)
	if err != nil {
		return nil, err
	}
	node.name = *name
	node.ip = ip
	node.uiPort = *uiPort
	node.gossipPort = port

	return node.Validate()
}

func (node *Node) Validate() (*Node, error) {
	if len(node.name) == 0 {
		return nil, errors.New("Node should have a name")
	}

	if node.gossipPort <= 5000 || node.uiPort <= 5000 {
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

func getPeers(s string) {

}
