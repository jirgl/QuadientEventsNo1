package api

import (
	"encoding/json"
	"fmt"

	m "github.com/jirgl/quadient-events-no1/model"
)

//GetTask func fetches task from API
func GetTask() m.Task {
	fmt.Println("get")
	resp, _ := callGet("task")
	fmt.Println(resp)
	var task m.Task
	json.Unmarshal(resp, &task)

	return task
}

//SubmitTask func submits the shortest path
func SubmitTask(id, answer string) {
	callPut("task/"+id, answer)
}
