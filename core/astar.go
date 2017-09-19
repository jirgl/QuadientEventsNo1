package core

import (
	"container/heap"

	m "github.com/jirgl/quadient-events-no1/model"
)

type parentNode struct {
	node      *node
	direction string
}

type node struct {
	position        m.Position
	originData      string
	path            []string
	cost            int
	parentDirection string

	open   bool
	closed bool
	eval   int
	index  int //TODO pro que
}

//Traveler interface gets the nearest nodes
type Traveler interface {
	getNextNodes(n *node) []*node
}

func findPath(start, end *node, traveler Traveler) []string {
	que := &priorityQueue{}
	heap.Init(que)
	start.open = true
	heap.Push(que, start)

	for {
		if que.Len() == 0 {
			return []string{}
		}

		current := heap.Pop(que).(*node)
		current.open = false
		current.closed = true

		if current.position == end.position {
			return current.path
		}

		for _, neighbor := range traveler.getNextNodes(current) {
			cost := current.cost + neighbor.cost
			if cost < neighbor.cost {
				if neighbor.open {
					heap.Remove(que, neighbor.index)
				}
				neighbor.open = false
				neighbor.closed = false
			}
			if !neighbor.open && !neighbor.closed {
				neighbor.cost = cost
				neighbor.open = true
				neighbor.eval = cost + getHeuristicEvaluation(neighbor, end)
				neighbor.path = append(current.path, neighbor.parentDirection)
				heap.Push(que, neighbor)
			}
		}
	}
}

func getHeuristicEvaluation(from, to *node) int {
	return 0
}

func reversePath(path []string) []string {
	last := len(path) - 1
	for i := 0; i < len(path)/2; i++ {
		path[i], path[last-i] = path[last-i], path[i]
	}

	return path
}
