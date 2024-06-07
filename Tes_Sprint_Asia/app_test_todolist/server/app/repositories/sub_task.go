package repositories

import (
	"errors"
	"todolist_sprint_asia/app/models"

	"gorm.io/gorm"
)

type SubTaskRepository struct {
	DB *gorm.DB
}

func GetSubTaskRepository(db *gorm.DB) *SubTaskRepository {
	return &SubTaskRepository{DB: db}
}

func (r *SubTaskRepository) GetSubTaskByID(id uint) (*models.SubTask, error) {
	var subTask models.SubTask

	result := r.DB.Where("status != ?", "deleted").First(&subTask, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("sub task not found")
		}

		return nil, result.Error
	}

	return &subTask, nil
}

func (r *SubTaskRepository) GetAllSubTaskByTaskID(taskID uint) ([]models.SubTask, error) {
	var subTasks []models.SubTask

	result := r.DB.Where("status != ? AND task_id = ?", "deleted", taskID).Find(&subTasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return subTasks, nil
}

func (r *SubTaskRepository) AddSubTask(subTask *models.SubTask) error {
	result := r.DB.Create(subTask)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SubTaskRepository) UpdateSubTask(subTask models.SubTask, id uint) error {
	result := r.DB.Where("id = ?", id).Updates(subTask)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SubTaskRepository) DeleteSubTask(id uint) error {
	result := r.DB.Delete(&models.SubTask{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("0")
	}

	return nil
}
