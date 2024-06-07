package models

type SubTask struct {
	Base

	TaskID *int `gorm:"foreignKey:TaskID"`
}

type SubTaskRequest struct {
	TaskID      *int
	Description *string
	Status      *string
}
