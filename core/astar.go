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
	regularScore    int
	heuristicScore  int
	totalScore      int
}

//Traveler interface gets the nearest nodes
type Traveler interface {
	getNextNodes(n *Node) []*Node
}

//FindPath func finds the shortest path
func FindPath(from, to *Node, traveler Traveler) []string {
	closedSet = map[m.Position]*Node{}
	openSet = map[m.Position]*Node{}
	openSet[from.Position] = from
	from.regularScore = 0
	from.heuristicScore = getHeuristicEvaluation(from, to)
	from.totalScore = from.heuristicScore

	for len(openSet) != 0 {
		x := getBestNode()
		if x.Position == to.Position {
			return createPath(x)
		}
		delete(openSet, x.Position)
		closedSet[x.Position] = x
		for _, y := range traveler.getNextNodes(x) {
			_, exists := closedSet[y.Position]
			if exists == true {
				continue
			}

			currentGScore := x.regularScore + y.regularScore
			currentIsBetter := false

			_, exists = openSet[y.Position]
			if exists != true {
				openSet[y.Position] = y
				currentIsBetter = true
			} else if currentGScore < y.regularScore {
				currentIsBetter = true
			} else {
				currentIsBetter = false
			}

			if currentIsBetter {
				y.parent = x
				y.regularScore = currentGScore
				y.heuristicScore = getHeuristicEvaluation(y, to)
				y.totalScore = y.regularScore + y.heuristicScore
			}
		}
	}

	return []string{}
}

func getBestNode() *Node {
	var best *Node
	for _, node := range openSet {
		if best == nil {
			best = node
		} else if node.totalScore < best.totalScore {
			best = node
		}
	}
	return best
}

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
