package models

type CreateTask struct {
	TaskID   string `form:"taskID" json:"taskID"`
	TaskName string `form:"taskName" json:"taskName"`
	TaskDesc string `form:"taskDescription" json:"taskDescription"`
}
