package core

import (
	"math"

	m "github.com/jirgl/quadient-events-no1/model"
)

//ArrayTraveler struct is layer for getting surrounding nodes from array
type ArrayTraveler struct {
	array   []string
	dimSize int
}

//Init func initializes traveler
func (at *ArrayTraveler) Init(array []string) {
	at.array = array
	at.dimSize = int(math.Sqrt(float64(len(array))))
}

func (at ArrayTraveler) getNextNodes(n *Node) []*Node {
	_, directions := m.ParseNode(n.OriginData)
	nodes := []*Node{}
	for _, direction := range directions {
		var x, y int
		if direction == 'L' {
			x = n.Position.X - 1
			y = n.Position.Y
		} else if direction == 'U' {
			x = n.Position.X
			y = n.Position.Y - 1
		} else if direction == 'R' {
			x = n.Position.X + 1
			y = n.Position.Y
		} else {
			x = n.Position.X
			y = n.Position.Y + 1
		}
		cost, _ := m.ParseNode(at.array[y*at.dimSize+x])
		nodes = append(nodes, &Node{
			OriginData:      at.array[y*at.dimSize+x],
			cost:            cost,
			parentDirection: string(direction),
			Position: m.Position{
				X: x,
				Y: y,
			},
		})
	}
	return nodes
}
