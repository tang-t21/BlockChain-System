package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Left     *Node
	Right    *Node
	Data     []byte
	NumChild int
}

func NewMerkleTree(data [][]byte) *Tree {
	if len(data) == 0 {
		return nil
	}
	var nodes []Node

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum) // the leaf nodes
		nodes = append(nodes, *node)
	}
	for true {
		if len(nodes) == 1 {
			break
		}
		var newLevel []Node

		for j := 0; j < len(nodes); j += 2 {
			if j == len(nodes)-1 {
				node := NewMerkleNode(&nodes[j], nil, nil)
				newLevel = append(newLevel, *node)
			} else {
				node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
				newLevel = append(newLevel, *node)
			}
		}
		nodes = newLevel
	}

	mTree := Tree{&nodes[0]}

	return &mTree
}

func NewMerkleNode(left, right *Node, data []byte) *Node {
	mNode := Node{}
	if left == nil && right == nil {
		// hash := sha256.Sum256(data)
		mNode.Data = data
		mNode.NumChild = 0
	} else if right == nil {
		children_data := left.Data
		hash := sha256.Sum256(children_data)
		mNode.Data = hash[:]
		mNode.NumChild = 1
	} else {
		children_data := append(left.Data, right.Data...)
		hash := sha256.Sum256(children_data)
		mNode.Data = hash[:]
		mNode.NumChild = 2
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}

func (t Tree) String() string {
	var lines []string
	if t.Root == nil {
		return "--- empty tree"
	}
	lines = append(lines, fmt.Sprintf("--- begin tree: %x", t.Root.Data[:2]))
	var nodes []Node
	nodes = append(nodes, *t.Root)
	for true {
		var data [][]byte
		var nextLevel []Node
		for _, node := range nodes {
			data = append(data, node.Data)
			if node.NumChild == 1 {
				nextLevel = append(nextLevel, *node.Left)
			} else if node.NumChild == 2 {
				nextLevel = append(nextLevel, *node.Left)
				nextLevel = append(nextLevel, *node.Right)
			}
		}
		var line []string
		for _, datum := range data {
			line = append(line, fmt.Sprintf("%x", datum[:2]))
		}
		lines = append(lines, strings.Join(line, "  "))

		if nodes[0].NumChild == 0 {
			break
		}
		nodes = nextLevel
	}
	return strings.Join(lines, "\n")
}
