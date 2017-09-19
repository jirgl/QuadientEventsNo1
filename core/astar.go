package core

import (
	"container/heap"

	m "github.com/jirgl/quadient-events-no1/model"
)

//Node struct
type Node struct {
	Position   m.Position
	OriginData string

	parent          *Node
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
	getNextNodes(n *Node) []*Node
}

//FindPath func finds the shortest path
func FindPath(from, to *Node, traveler Traveler) []string {
	que := &priorityQueue{}
	heap.Init(que)
	from.open = true
	heap.Push(que, from)

	for {
		if que.Len() == 0 {
			return []string{}
		}

		current := heap.Pop(que).(*Node)
		current.open = false
		current.closed = true

		if current.Position == to.Position {
			return current.path
		}

		for _, neighbor := range traveler.getNextNodes(current) {
			if current.parent != nil && current.parent.Position == neighbor.Position {
				continue
			}
			cost := current.cost + neighbor.cost
			if cost < neighbor.cost {
				if neighbor.open {
					heap.Remove(que, neighbor.index)
				}
				neighbor.open = false
				neighbor.closed = false
			}
			if !neighbor.open && !neighbor.closed {
				neighbor.open = true
				neighbor.cost = cost
				neighbor.eval = cost + getHeuristicEvaluation(neighbor, to)
				neighbor.path = append(current.path, neighbor.parentDirection)
				neighbor.parent = current
				heap.Push(que, neighbor)
			}
		}
	}
}

func getHeuristicEvaluation(from, to *Node) int {
	return 0
}
