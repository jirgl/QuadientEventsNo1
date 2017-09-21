package core

import (
	m "github.com/jirgl/quadient-events-no1/model"
)

var openSet map[m.Position]*Node
var closedSet map[m.Position]*Node

//Node struct
type Node struct {
	OriginData string
	Position   m.Position

	parentDirection string
	parent          *Node

	regularScore   int
	heuristicScore int
	totalScore     int
}

//Traveler interface gets the nearest nodes
type Traveler interface {
	getNextNodes(n *Node) []*Node
}

//FindPath func finds the shortest path
func FindPath(from, to *Node, traveler Traveler) []string {
	heap := InitFibHeap()
	closedSet = map[m.Position]*Node{}
	openSet = map[m.Position]*Node{}
	openSet[from.Position] = from
	heap.Insert(0, from)
	from.regularScore = 0
	from.heuristicScore = getHeuristicEvaluation(from, to)
	from.totalScore = from.heuristicScore

	for len(openSet) != 0 {
		current := getBestNode(heap)
		if current.Position == to.Position {
			return createPath(current)
		}
		delete(openSet, current.Position)
		closedSet[current.Position] = current
		for _, neighbor := range traveler.getNextNodes(current) {
			_, exists := closedSet[neighbor.Position]
			if exists == true {
				continue
			}

			currentRegularScore := current.regularScore + neighbor.regularScore
			currentIsBetter := false

			_, exists = openSet[neighbor.Position]
			if exists != true {
				heap.Insert(float64(currentRegularScore), neighbor)
				openSet[neighbor.Position] = neighbor
				currentIsBetter = true
			} else if currentRegularScore < neighbor.regularScore {
				currentIsBetter = true
			} else {
				currentIsBetter = false
			}

			if currentIsBetter {
				neighbor.parent = current
				neighbor.regularScore = currentRegularScore
				neighbor.heuristicScore = getHeuristicEvaluation(neighbor, to)
				neighbor.totalScore = neighbor.regularScore + neighbor.heuristicScore
			}
		}
	}

	return []string{}
}

func getBestNode(fh *FibHeap) *Node {
	_, min := fh.ExtractMin()
	return min.(*Node)
}

//Manhattan distance
func getHeuristicEvaluation(from, to *Node) int {
	absX := from.Position.X - to.Position.X
	if absX < 0 {
		absX = -absX
	}
	absY := from.Position.Y - to.Position.Y
	if absY < 0 {
		absY = -absY
	}
	r := absX + absY

	return r
}

func createPath(n *Node) []string {
	path := []string{}
	for n.parentDirection != "" {
		path = append([]string{n.parentDirection}, path...)
		n = n.parent
	}

	return path
}
