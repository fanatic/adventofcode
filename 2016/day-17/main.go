package main

import (
	"crypto/md5"
	"fmt"
)

//var passcode = "hijkl"
//var passcode = "ihgpwlah"

//var passcode = "kglvqrro"
//var passcode = "ulqzkmiv"

var passcode = "hhhxzeay"

func main() {
	queue := []Node{}
	longestPath := 0

	// Start at root node
	queue = append(queue, Node{1, 1, ""})

	// While queue is not empty
	//OUTER:
	for 0 < len(queue) {
		// Dequeue current node
		node := queue[0]
		queue = queue[1:]
		//fmt.Printf("Checking: %v.  Rest: %v\n", node, queue)

		// Check for success condition
		if node.x == 4 && node.y == 4 {
			fmt.Printf("Len: %d, MaxLen: %d, Queue: %d\n", len(node.path), longestPath, len(queue))

			// Part 2:
			if len(node.path) > longestPath {
				longestPath = len(node.path)
			}
			continue
		}

		// For each node that's adjacent to current
		for _, n := range node.adjacent() {

			// Add to queue
			queue = append(queue, n)
		}
	}

	fmt.Printf("Longest path: %d\n", longestPath)
}

type Node struct {
	x    int
	y    int
	path string
}

func (n Node) adjacent() []Node {
	adj := []Node{}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(passcode+n.path)))
	// doors 0=up, 1=down, 2=left, and 3=righ
	//fmt.Printf("Input: %s, Hash: %s\n", passcode+n.path, hash[0:4])

	if n.x < 4 && isOpen(hash[3]) {
		adj = append(adj, Node{n.x + 1, n.y, n.path + "R"})
	}
	if n.y < 4 && isOpen(hash[1]) {
		adj = append(adj, Node{n.x, n.y + 1, n.path + "D"})
	}
	if n.x > 1 && isOpen(hash[2]) {
		adj = append(adj, Node{n.x - 1, n.y, n.path + "L"})
	}
	if n.y > 1 && isOpen(hash[0]) {
		adj = append(adj, Node{n.x, n.y - 1, n.path + "U"})
	}
	return adj
}

func isOpen(c byte) bool {
	if c == 'b' || c == 'c' || c == 'd' || c == 'e' || c == 'f' {
		return true
	}
	return false
}
