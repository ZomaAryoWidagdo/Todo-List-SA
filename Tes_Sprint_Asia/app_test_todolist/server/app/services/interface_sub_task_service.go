package services

import "todolist_sprint_asia/app/models"

type SubTaskServiceInterface interface {
	GetSubTaskByID(id uint) (*models.SubTask, error)
	GetAllSubTaskByTaskID(taskID uint) ([]models.SubTask, error)
	AddSubTask(subTaskRequest models.SubTaskRequest) (*models.SubTask, error)
	UpdateSubTask(subTaskRequest models.SubTaskRequest, id uint) (*models.SubTask, error)
	DeleteSubTask(id uint) error
}
