package services

import (
	"todolist_sprint_asia/app/models"
	"todolist_sprint_asia/app/models/enum"
	"todolist_sprint_asia/app/repositories"
	"todolist_sprint_asia/app/utils"
)

type TaskService struct {
	TaskRepo repositories.TaskRepositoryInterface
}

func GetTaskService(taskrepo repositories.TaskRepositoryInterface) *TaskService {
	return &TaskService{TaskRepo: taskrepo}
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	task, err := s.TaskRepo.GetTaskByID(id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetAllTask(status string) ([]models.Task, error) {
	tasks, err := s.TaskRepo.GetAllTask(status)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) AddTask(taskRequest models.TaskRequest) (*models.Task, error) {
	deadline := utils.StringToTime(taskRequest.Deadline)

	task := models.Task{
		Base: models.Base{
			Description: taskRequest.Description,
		},
		Deadline: deadline,
		SubTasks: taskRequest.SubTask,
	}

	if err := s.TaskRepo.AddTask(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *TaskService) UpdateTask(taskRequest models.TaskRequest, id uint) (*models.Task, error) {

	deadline := utils.StringToTime(taskRequest.Deadline)

	task := models.Task{
		Base: models.Base{
			Description: taskRequest.Description,
			Status:      taskRequest.Status,
		},
		Deadline: deadline,
	}

	if err := s.TaskRepo.UpdateTask(task, id); err != nil {
		return nil, err
	}

	updatedTask, err := s.TaskRepo.GetTaskByID(id)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *TaskService) DeleteTask(id uint) error {

	setTaskStatus := models.Task{
		Base: models.Base{
			Status: &enum.Deleted,
		},
	}

	if err := s.TaskRepo.UpdateTask(setTaskStatus, id); err != nil {
		return err
	}

	err := s.TaskRepo.DeleteTask(id)

	if err != nil {
		return err
	}

	return nil
}
