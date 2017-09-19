package main

import (
	"fmt"
	"strings"

	api "github.com/jirgl/quadient-events-no1/api"
	core "github.com/jirgl/quadient-events-no1/core"
	m "github.com/jirgl/quadient-events-no1/model"
)

func main() {
	task := api.GetTask()
	fmt.Println(task)
	traveler := core.ArrayTraveler{}
	traveler.Init(task.Map.Areas)
	path := core.FindPath(&core.Node{
		OriginData: "5-RD",
		Position: m.Position{
			X: 0,
			Y: 0,
		},
	}, &core.Node{
		OriginData: "5-RD",
		Position: m.Position{
			X: 0,
			Y: 0,
		},
	}, traveler)

	fmt.Println(path)
	api.SubmitTask(task.ID, strings.Join(path, ""))
}
