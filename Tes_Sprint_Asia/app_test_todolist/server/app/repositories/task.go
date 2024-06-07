package repositories

import (
	"errors"
	"todolist_sprint_asia/app/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func GetTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task

	result := r.DB.Where("status != ?", "deleted").Preload("SubTasks", "status != ?", "deleted").First(&task, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}

		return nil, result.Error
	}

	return &task, nil
}

func (r *TaskRepository) GetAllTask(status string) ([]models.Task, error) {
	var tasks []models.Task

	result := r.DB.Where("status = ?", status).Preload("SubTasks", func(db *gorm.DB) *gorm.DB {
		return db.Where("status != ?", "deleted").Order("id").Order("status")
	}).
		Order("id DESC").Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *TaskRepository) AddTask(task *models.Task) error {
	result := r.DB.Create(task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepository) UpdateTask(task models.Task, id uint) error {
	result := r.DB.Where("id = ?", id).Updates(task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepository) DeleteTask(id uint) error {
	result := r.DB.Delete(&models.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("0")
	}

	return nil
}
