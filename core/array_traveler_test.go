package core

import (
	"testing"

	m "github.com/jirgl/quadient-events-no1/model"
	"github.com/stretchr/testify/assert"
)

func TestGettingCorrectNodeFromArray(t *testing.T) {
	traveler := ArrayTraveler{}
	traveler.Init([]string{"5-R", "1-RDL", "10-DL", "2-RD", "1-UL", "1-UD", "2-RU", "1-RL", "2-UL"})

	nodes := traveler.getNextNodes(&Node{
		OriginData: "1-RDL",
		Position:   m.Position{X: 1, Y: 0},
	})
	assert.Equal(t, &Node{OriginData: "10-DL", parentDirection: "R", cost: 10, Position: m.Position{X: 2, Y: 0}}, nodes[0])
	assert.Equal(t, &Node{OriginData: "1-UL", parentDirection: "D", cost: 1, Position: m.Position{X: 1, Y: 1}}, nodes[1])
	assert.Equal(t, &Node{OriginData: "5-R", parentDirection: "L", cost: 5, Position: m.Position{X: 0, Y: 0}}, nodes[2])

	nodes = traveler.getNextNodes(&Node{
		OriginData: "2-RD",
		Position:   m.Position{X: 1, Y: 1},
	})
	assert.Equal(t, &Node{OriginData: "1-UD", parentDirection: "R", cost: 1, Position: m.Position{X: 2, Y: 1}}, nodes[0])
	assert.Equal(t, &Node{OriginData: "1-RL", parentDirection: "D", cost: 1, Position: m.Position{X: 1, Y: 2}}, nodes[1])
}
