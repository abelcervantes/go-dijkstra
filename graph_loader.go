package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const(
	file = "graph.json"
)

type graph map[uint]map[uint]uint

func loadGraph() (graph, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var g graph
	if err := json.Unmarshal(bytes, &g); err != nil {
		return nil, err
	}

	return g, nil
}

func initQueue(g graph, source uint) (map[uint]*node, PriorityQueue, error) {
	nodes := make(map[uint]*node, 0)
	for n, neighbours := range g {
		nodes[n] = &node{
			id:         n,
			distance:   ^uint(0),
			neighbours: make([]*node, len(neighbours)),
		}
	}

	for n, neighbours := range g {
		i := 0
		for neigh := range neighbours {
			nodes[n].neighbours[i] = nodes[neigh]
			i++
		}
	}

	if _, exist := nodes[source]; !exist {
		return nil, nil, fmt.Errorf("invalid source: %d", source)
	}

	nodes[source].distance = 0

	pq := make(PriorityQueue, len(nodes))

	i := 0
	for k := range nodes {
		nodes[k].index = i
		pq[i] = nodes[k]
		i++
	}
	heap.Init(&pq)

	return nodes, pq, nil
}
