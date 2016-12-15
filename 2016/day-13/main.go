package main

import "fmt"

const (
	passable = iota
	blocked
	path
)

func main() {
	maze := &M{
		MagicNum: 1362,
		Cache:    map[Node]int{},
	}
	size := 40
	goal := Node{31, 39}

	maze.Print(size)
	maze.findShortestPath(goal)
	maze.Print(size)
}

// M is Maze Cache
type M struct {
	MagicNum int
	Cache    map[Node]int
}

func (m *M) Get(x, y int) int {
	n, exists := m.Cache[Node{x, y}]
	if exists {
		return n
	}
	sum := x*x + 3*x + 2*x*y + y + y*y + m.MagicNum
	bSum := fmt.Sprintf("%b", sum)
	numOnes := 0
	for _, c := range bSum {
		if c == '1' {
			numOnes++
		}
	}
	n = numOnes % 2
	m.Cache[Node{x, y}] = n
	return n
}

func (m *M) MarkPath(x, y int) {
	m.Cache[Node{x, y}] = path
}

func (m *M) Print(size int) {
	fmt.Printf("  0123456789\n")
	for y := 0; y < size; y++ {
		fmt.Printf("%d ", y)
		for x := 0; x < size; x++ {
			if m.Get(x, y) == passable {
				fmt.Printf(".")
			} else if m.Get(x, y) == blocked {
				fmt.Printf("#")
			} else if m.Get(x, y) == path {
				fmt.Printf("O")
			}
		}
		fmt.Printf("\n")
	}
}

func (maze *M) findShortestPath(successNode Node) {
	queue := []Node{}
	distances := map[Node]int{}
	parents := map[Node]Node{}

	// Start at root node
	queue = append(queue, Node{1, 1})
	distances[Node{1, 1}] = 0

	// While queue is not empty
OUTER:
	for 0 < len(queue) {
		// Dequeue current node
		node := queue[0]
		queue = queue[1:]
		//fmt.Printf("Checking: %v.  Rest: %v\n", node, queue)

		// For each node that's adjacent to current
		for _, n := range node.adjacent() {
			// Check if we've seen it before
			_, exists := distances[n]
			if !exists {
				// Check for blockages
				if maze.Get(n.x, n.y) == blocked {
					continue
				}

				// Update our statistics
				distances[n] = distances[node] + 1
				parents[n] = node

				// Check for succes condition
				// if n.x == successNode.x && n.y == successNode.y {
				// 	fmt.Printf("Success. Distance: %d\n", distances[n])
				// 	break OUTER
				// }
				// Add to queue
				queue = append(queue, n)
			}
		}

		// Check for second success condition
		if distances[node] == 50 {
			fmt.Printf("Success. Nodes reached: %d\n", len(distances))
			break OUTER
		}
	}

	// Now mark our path
	for n := successNode; ; {
		maze.MarkPath(n.x, n.y)
		p, exists := parents[n]
		if !exists {
			break
		}
		n = p
	}
}

type Node struct {
	x int
	y int
}

func (n Node) adjacent() []Node {
	adj := []Node{
		Node{n.x + 1, n.y},
		Node{n.x, n.y + 1},
	}
	if n.x > 0 {
		adj = append(adj, Node{n.x - 1, n.y})
	}
	if n.y > 0 {
		adj = append(adj, Node{n.x, n.y - 1})
	}
	return adj
}
