package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	id         uint
	distance   uint
	neighbours []*node
	prevID     uint
	index      int
}

func (n node) String() string {
	return fmt.Sprintf("\nid: %d\tdistance: %d\tprevious: %d\n", n.id, n.distance, n.prevID)
}

func initGraph() (map[uint]*node, map[uint]map[uint]uint) {
	n1 := node{
		id:       1,
		distance: 0,
	}

	n2 := node{
		id:       2,
		distance: ^uint(0),
	}

	n3 := node{
		id:       3,
		distance: ^uint(0),
	}

	n1.neighbours = []*node{&n2, &n3}
	n2.neighbours = []*node{&n3}

	nodes := map[uint]*node{
		n1.id: &n1,
		n2.id: &n2,
		n3.id: &n3,
	}

	costMatrix := map[uint]map[uint]uint{
		1: {
			2: 5,
			3: 10,
		},
		2: {
			3: 2,
		},
	}

	return nodes, costMatrix
}

func printPaths(nodes map[uint]*node) {
	fmt.Println("\n...shortest paths...")

	source := nodes[1]
	for _, v := range nodes {
		if v.id == source.id {
			continue
		}

		path := make([]uint, 0)
		path = append(path, source.id)
		prev := v.prevID
		for prev != source.id {
			path = append(path, nodes[prev].id)
			prev = nodes[prev].prevID
		}
		path = append(path, v.id)
		fmt.Printf("from %d to %d: %v\n", source.id, v.id, path)
	}
}

func main() {
	nodes, weights := initGraph()

	pq := make(PriorityQueue, len(nodes))

	i := 0
	for k := range nodes {
		nodes[k].index = i
		pq[i] = nodes[k]
		i++
	}
	heap.Init(&pq)

	fmt.Println(pq)
	fmt.Println(weights)

	for len(pq) > 0 {
		current := heap.Pop(&pq).(*node)
		for neighIndex := range current.neighbours {
			tmp := current.distance + weights[current.id][current.neighbours[neighIndex].id]
			if tmp < current.neighbours[neighIndex].distance {
				current.neighbours[neighIndex].distance = tmp
				current.neighbours[neighIndex].prevID = current.id
				heap.Fix(&pq, current.neighbours[neighIndex].index)
			}
		}
	}

	printPaths(nodes)
}
