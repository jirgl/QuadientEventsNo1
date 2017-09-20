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
	assert.Equal(t, &Node{OriginData: "10-DL", regularScore: 10, parentDirection: "R", Position: m.Position{X: 2, Y: 0}}, nodes[0])
	assert.Equal(t, &Node{OriginData: "1-UL", regularScore: 1, parentDirection: "D", Position: m.Position{X: 1, Y: 1}}, nodes[1])
	assert.Equal(t, &Node{OriginData: "5-R", regularScore: 5, parentDirection: "L", Position: m.Position{X: 0, Y: 0}}, nodes[2])

	nodes = traveler.getNextNodes(&Node{
		OriginData: "2-RD",
		Position:   m.Position{X: 1, Y: 1},
	})
	assert.Equal(t, &Node{OriginData: "1-UD", regularScore: 1, parentDirection: "R", Position: m.Position{X: 2, Y: 1}}, nodes[0])
	assert.Equal(t, &Node{OriginData: "1-RL", regularScore: 1, parentDirection: "D", Position: m.Position{X: 1, Y: 2}}, nodes[1])
}
