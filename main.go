package main

import (
	"fmt"
	"strings"
	"time"

	api "github.com/jirgl/quadient-events-no1/api"
	core "github.com/jirgl/quadient-events-no1/core"
)

func main() {
	task := api.GetTask()
	start := time.Now().UnixNano()
	traveler := core.ArrayTraveler{}
	traveler.Init(task.Map.Areas)
	fmt.Println("start finding")
	path := core.FindPath(
		traveler.GetNode(task.Astroants.X, task.Astroants.Y),
		traveler.GetNode(task.Sugar.X, task.Sugar.Y),
		traveler)

	fmt.Println("done in ", (time.Now().UnixNano()-start)/1000)
	api.SubmitTask(task.ID, strings.Join(path, ""))
}
