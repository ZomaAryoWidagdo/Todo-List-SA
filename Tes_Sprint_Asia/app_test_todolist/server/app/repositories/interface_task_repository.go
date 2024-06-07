package repositories

import "todolist_sprint_asia/app/models"

type TaskRepositoryInterface interface {
	GetTaskByID(id uint) (*models.Task, error)
	GetAllTask(status string) ([]models.Task, error)
	AddTask(task *models.Task) error
	UpdateTask(task models.Task, id uint) error
	DeleteTask(id uint) error
}
