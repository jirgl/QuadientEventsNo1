package core

import (
	"testing"

	m "github.com/jirgl/quadient-events-no1/model"
	"github.com/stretchr/testify/assert"
)

func TestGettingCorrectNodeFromArray(t *testing.T) {
	traveler := arrayTraveler{}
	traveler.init([]string{"5-R", "1-RDL", "10-DL", "2-RD", "1-UL", "1-UD", "2-RU", "1-RL", "2-UL"})

	nodes := traveler.getNextNodes(&node{
		originData: "1-RDL",
		position:   m.Position{X: 1, Y: 0},
	})
	assert.Equal(t, &node{originData: "10-DL", cost: 10, position: m.Position{X: 2, Y: 0}}, nodes[0])
	assert.Equal(t, &node{originData: "1-UL", cost: 1, position: m.Position{X: 1, Y: 1}}, nodes[1])
	assert.Equal(t, &node{originData: "5-R", cost: 5, position: m.Position{X: 0, Y: 0}}, nodes[2])

	nodes = traveler.getNextNodes(&node{
		originData: "2-RD",
		position:   m.Position{X: 1, Y: 1},
	})
	assert.Equal(t, &node{originData: "1-UD", cost: 1, position: m.Position{X: 2, Y: 1}}, nodes[0])
	assert.Equal(t, &node{originData: "1-RL", cost: 1, position: m.Position{X: 1, Y: 2}}, nodes[1])
}
