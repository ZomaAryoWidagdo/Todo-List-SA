package models

import "time"

type Task struct {
	Base

	Deadline *time.Time

	SubTasks []SubTask `json:"subTasks,omitempty"`
}

type TaskRequest struct {
	Description *string
	Status      *string
	Deadline    *string

	SubTask []SubTask `json:"subTasks,omitempty"`
}
