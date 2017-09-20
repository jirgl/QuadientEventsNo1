package model

import (
	"strconv"
	"strings"
)

//Map struct contains all area's nodes
type Map struct {
	Areas []string `json:"areas"`
}

//ParseNode func parses node and returns evaluation and next directions
func ParseNode(node string) (int, string) {
	parts := strings.Split(node, "-")
	evaluation, _ := strconv.Atoi(parts[0])
	return evaluation, parts[1]
}

//Position struct contains node position
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//Task struct
type Task struct {
	ID               string   `json:"id"`
	StartedTimestamp int64    `json:"startedTimestamp"`
	Map              Map      `json:"map"`
	Astroants        Position `json:"astroants"`
	Sugar            Position `json:"sugar"`
}

//SubmitTask struct
type SubmitTask struct {
	Path string `json:"path"`
}
