package core

import (
	"math"

	m "github.com/jirgl/quadient-events-no1/model"
)

type arrayTraveler struct {
	array   []string
	dimSize int
}

func (at *arrayTraveler) init(array []string) {
	at.array = array
	at.dimSize = int(math.Sqrt(float64(len(array))))
}

func (at arrayTraveler) getNextNodes(n *node) []*node {
	_, directions := m.ParseNode(n.originData)
	nodes := []*node{}
	for _, direction := range directions {
		var x, y int
		if direction == 'L' {
			x = n.position.X - 1
			y = n.position.Y
		} else if direction == 'U' {
			x = n.position.X
			y = n.position.Y - 1
		} else if direction == 'R' {
			x = n.position.X + 1
			y = n.position.Y
		} else {
			x = n.position.X
			y = n.position.Y + 1
		}
		cost, _ := m.ParseNode(at.array[y*at.dimSize+x])
		nodes = append(nodes, &node{
			originData:      at.array[y*at.dimSize+x],
			cost:            cost,
			parentDirection: string(direction),
			position: m.Position{
				X: x,
				Y: y,
			},
		})
	}
	return nodes
}
