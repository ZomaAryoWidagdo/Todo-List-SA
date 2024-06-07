package repositories

import "todolist_sprint_asia/app/models"

type SubTaskRepositoryInterface interface {
	GetSubTaskByID(id uint) (*models.SubTask, error)
	GetAllSubTaskByTaskID(taskID uint) ([]models.SubTask, error)
	AddSubTask(subTask *models.SubTask) error
	UpdateSubTask(subTask models.SubTask, id uint) error
	DeleteSubTask(id uint) error
}
