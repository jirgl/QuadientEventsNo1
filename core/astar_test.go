package core

import (
	"testing"

	m "github.com/jirgl/quadient-events-no1/model"
	"github.com/stretchr/testify/assert"
)

func TestFindingEndInMinimalCase(t *testing.T) {
	areas := []string{"5-RD", "1-LD", "2-UR", "5-LU"}
	traveler := ArrayTraveler{}
	traveler.Init(areas)
	path := FindPath(traveler.GetNode(0, 0), traveler.GetNode(1, 1), traveler)

	assert.Equal(t, []string{"R", "D"}, path)
}

func TestFindingPathInExample(t *testing.T) {
	task := m.Task{
		ID:               "2727",
		StartedTimestamp: 1503929807498,
		Map: m.Map{
			Areas: []string{"5-R", "1-RDL", "10-DL", "2-RD", "1-UL", "1-UD", "2-RU", "1-RL", "2-UL"},
		},
		Astroants: m.Position{X: 1, Y: 0},
		Sugar:     m.Position{X: 2, Y: 1},
	}

	traveler := ArrayTraveler{}
	traveler.Init(task.Map.Areas)
	path := FindPath(traveler.GetNode(1, 0), traveler.GetNode(2, 1), traveler)

	assert.Equal(t, []string{"D", "L", "D", "R", "R", "U"}, path)
}
