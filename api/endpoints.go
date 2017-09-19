package api

import (
	"encoding/json"

	m "github.com/jirgl/quadient-events-no1/model"
)

//GetTask func fetches task from API
func GetTask() {
	resp, _ := callGet("task")
	var task m.Task
	json.Unmarshal(resp, &task)
}

//SubmitTask func submits the shortest path
func SubmitTask() {
	callPut()
}
