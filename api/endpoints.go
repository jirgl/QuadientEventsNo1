package api

import (
	"encoding/json"
	"fmt"

	m "github.com/jirgl/quadient-events-no1/model"
)

/*
endpoints.go provides endpoints of server
*/

//GetTask func fetches task from API
func GetTask() m.Task {
	response, _ := callGet("task")
	var task m.Task
	json.Unmarshal(response, &task)

	return task
}

//SubmitTask func submits the shortest path
func SubmitTask(id, answer string) {
	submitTask := m.SubmitTask{Path: answer}
	data, _ := json.Marshal(submitTask)
	response, _ := callPut("task/"+id, data)
	fmt.Println(string(response))
}
