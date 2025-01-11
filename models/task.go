package models

type TaskJson struct {
	TaskID   int64
	TaskName string `json:"taskName"`
	TaskDesc string `json:"taskDescription"`
}
