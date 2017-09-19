package core

import (
	"testing"

	m "github.com/jirgl/quadient-events-no1/model"
	"github.com/stretchr/testify/assert"
)

func TestFindEndInMinimalCase(t *testing.T) {
	areas := []string{"5-RD", "1-LD", "2-UR", "5-LU"}
	travaler := ArrayTraveler{}
	travaler.Init(areas)
	path := FindPath(&Node{
		OriginData: "5-RD",
		Position: m.Position{
			X: 0,
			Y: 0,
		},
	}, &Node{
		OriginData: "5-LU",
		Position: m.Position{
			X: 1,
			Y: 1,
		},
	}, travaler)

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

	travaler := ArrayTraveler{}
	travaler.Init(task.Map.Areas)
	path := FindPath(&Node{
		OriginData: "1-RDL",
		Position: m.Position{
			X: task.Astroants.X,
			Y: task.Astroants.Y,
		},
	}, &Node{
		Position: m.Position{
			X: task.Sugar.X,
			Y: task.Sugar.Y,
		},
	}, travaler)

	assert.Equal(t, []string{"D", "L", "D", "R", "R", "U"}, path)
}
