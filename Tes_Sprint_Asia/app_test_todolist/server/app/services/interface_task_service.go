package services

import "todolist_sprint_asia/app/models"

type TaskServiceInterface interface {
	GetTaskByID(id uint) (*models.Task, error)
	GetAllTask(status string) ([]models.Task, error)
	AddTask(taskRequest models.TaskRequest) (*models.Task, error)
	UpdateTask(taskRequest models.TaskRequest, id uint) (*models.Task, error)
	DeleteTask(id uint) error
}
