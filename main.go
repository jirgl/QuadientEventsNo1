package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	api "github.com/jirgl/quadient-events-no1/api"
	core "github.com/jirgl/quadient-events-no1/core"
	visual "github.com/jirgl/quadient-events-no1/visualization"
)

func runTask() int64 {
	task := api.GetTask()
	start := time.Now().UnixNano()
	traveler := core.ArrayTraveler{}
	traveler.Init(task.Map.Areas)
	fmt.Println("start finding")
	path := core.FindPath(
		traveler.GetNode(task.Astroants.X, task.Astroants.Y),
		traveler.GetNode(task.Sugar.X, task.Sugar.Y),
		traveler)

	duration := (time.Now().UnixNano() - start) / 1000000
	fmt.Println("done in", duration, "ms")
	api.SubmitTask(task.ID, strings.Join(path, ""))

	if len(os.Args) == 2 && os.Args[1] == "-v" {
		fmt.Println("Visualization started")
		visual.Visualize(task.ID, &traveler)
		fmt.Println("Visualization done")
	}

	return duration
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-d" {
		attempts := 10
		totalTime := int64(0)
		for i := 0; i < attempts; i++ {
			totalTime += runTask()
			time.Sleep(time.Second * 1)
		}
		fmt.Println("Benchmark time:", totalTime/10, "ms")
	} else {
		runTask()
	}
}
