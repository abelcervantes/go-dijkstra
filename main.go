package main

import (
	"container/heap"
	"fmt"
	"log"
)

type node struct {
	id         uint
	distance   uint
	neighbours []*node
	prevID     uint
	index      int
}

func (n node) String() string {
	neighs := make([]uint, len(n.neighbours))
	for i, neigh := range n.neighbours {
		neighs[i] = neigh.id
	}

	return fmt.Sprintf("\nid: %d\tindex: %d\tdistance: %d\tprevious: %d\tneighbours: %v\n",
		n.id, n.index, n.distance, n.prevID, neighs)
}

func dijkstraSearch(pq PriorityQueue, g graph) {
	for len(pq) > 0 {
		current := heap.Pop(&pq).(*node)
		for neighIndex := range current.neighbours {
			tmp := current.distance + g[current.id][current.neighbours[neighIndex].id]
			if tmp < current.neighbours[neighIndex].distance {
				current.neighbours[neighIndex].distance = tmp
				current.neighbours[neighIndex].prevID = current.id
				heap.Fix(&pq, current.neighbours[neighIndex].index)
			}
		}
	}
}

func printPaths(nodes map[uint]*node, source uint) {
	fmt.Println("\n...shortest paths...")

	sourceNode := nodes[source]
	for _, v := range nodes {
		if v.id == sourceNode.id {
			continue
		}

		path := make([]uint, 0)
		path = append(path, v.id)

		prev := v.prevID
		for prev != sourceNode.id {
			path = append(path, nodes[prev].id)
			prev = nodes[prev].prevID
		}

		path = append(path, sourceNode.id)

		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}

		fmt.Printf("from %d to %d: %v total distance: %d\n", sourceNode.id, v.id, path, v.distance)
	}
}

func main() {
	g, err := loadGraph()
	if err != nil {
		log.Fatal(err)
	}

	const source = 1

	nodes, pq, err := initQueue(g, source)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nodes)

	dijkstraSearch(pq, g)

	fmt.Println(nodes)

	printPaths(nodes, source)
}
