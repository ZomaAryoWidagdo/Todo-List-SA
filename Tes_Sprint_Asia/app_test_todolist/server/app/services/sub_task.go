package services

import (
	"todolist_sprint_asia/app/models"
	"todolist_sprint_asia/app/models/enum"
	"todolist_sprint_asia/app/repositories"
)

type SubTaskService struct {
	SubTaskRepo repositories.SubTaskRepositoryInterface
}

func GetSubTaskService(subTaskRepo repositories.SubTaskRepositoryInterface) *SubTaskService {
	return &SubTaskService{SubTaskRepo: subTaskRepo}
}

func (s *SubTaskService) GetSubTaskByID(id uint) (*models.SubTask, error) {
	subTask, err := s.SubTaskRepo.GetSubTaskByID(id)

	if err != nil {
		return nil, err
	}

	return subTask, nil
}

func (s *SubTaskService) GetAllSubTaskByTaskID(taskID uint) ([]models.SubTask, error) {
	subTasks, err := s.SubTaskRepo.GetAllSubTaskByTaskID(taskID)

	if err != nil {
		return nil, err
	}

	return subTasks, nil
}

func (s *SubTaskService) AddSubTask(subTaskRequest models.SubTaskRequest) (*models.SubTask, error) {
	subTask := models.SubTask{
		Base: models.Base{
			Description: subTaskRequest.Description,
			Status:      subTaskRequest.Status,
		},
		TaskID: subTaskRequest.TaskID,
	}

	if err := s.SubTaskRepo.AddSubTask(&subTask); err != nil {
		return nil, err
	}

	return &subTask, nil
}

func (s *SubTaskService) UpdateSubTask(subTaskRequest models.SubTaskRequest, id uint) (*models.SubTask, error) {
	subTask := models.SubTask{
		Base: models.Base{
			Description: subTaskRequest.Description,
			Status:      subTaskRequest.Status,
		},
	}

	if err := s.SubTaskRepo.UpdateSubTask(subTask, id); err != nil {
		return nil, err
	}

	updatedTask, err := s.SubTaskRepo.GetSubTaskByID(id)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *SubTaskService) DeleteSubTask(id uint) error {
	setSubTaskStatus := models.SubTask{
		Base: models.Base{
			Status: &enum.Deleted,
		},
	}

	if err := s.SubTaskRepo.UpdateSubTask(setSubTaskStatus, id); err != nil {
		return err
	}

	err := s.SubTaskRepo.DeleteSubTask(id)

	if err != nil {
		return err
	}

	return nil
}
