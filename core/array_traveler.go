package core

import (
	"math"

	m "github.com/jirgl/quadient-events-no1/model"
)

/*
array_traveler.go is interlayer between server data and algorithm needs
*/

//ArrayTraveler struct is layer for getting surrounding nodes from array
type ArrayTraveler struct {
	array   []string
	dimSize int
}

//Init func initializes traveler and calculate dimension size
func (at *ArrayTraveler) Init(array []string) {
	at.array = array
	at.dimSize = int(math.Sqrt(float64(len(array))))
}

//GetNode func returns node on specific position
func (at *ArrayTraveler) GetNode(x, y int) *Node {
	cost, _ := m.ParseNode(at.array[y*at.dimSize+x])
	return &Node{
		OriginData:   at.array[y*at.dimSize+x],
		regularScore: cost,
		Position: m.Position{
			X: x,
			Y: y,
		},
	}
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
			regularScore:    cost,
			parentDirection: string(direction),
			Position: m.Position{
				X: x,
				Y: y,
			},
		})
	}
	return nodes
}
