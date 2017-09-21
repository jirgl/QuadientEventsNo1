package visualization

import (
	"testing"

	core "github.com/jirgl/quadient-events-no1/core"
)

func TestVizualize(t *testing.T) {
	traveler := core.ArrayTraveler{}
	traveler.Init([]string{"5-R", "1-RDL", "10-DL", "2-RD", "1-UL", "1-UD", "2-RU", "1-RL", "2-UL"})
	core.FindPath(traveler.GetNode(1, 0), traveler.GetNode(2, 1), traveler)
	Visualize("test", &traveler)
}
