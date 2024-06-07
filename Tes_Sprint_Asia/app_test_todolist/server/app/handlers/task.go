package handlers

import (
	"net/http"
	"todolist_sprint_asia/app/models"
	"todolist_sprint_asia/app/services"
	"todolist_sprint_asia/app/utils"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	TaskService services.TaskServiceInterface
}

func GetTaskHandler(taskService services.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{TaskService: taskService}
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	taskID := c.Param("id")

	uintTaskID, err := utils.StringToUint(taskID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to convert ID type"})
	}

	task, err := h.TaskService.GetTaskByID(uintTaskID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get task"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetAllTask(c echo.Context) error {
	status := c.QueryParam("status")

	tasks, err := h.TaskService.GetAllTask(status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get all task"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) AddTask(c echo.Context) error {
	var taskRequest models.TaskRequest

	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if taskRequest.Description == nil || *taskRequest.Description == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "description cannot be empty"})
	}

	createdTask, err := h.TaskService.AddTask(taskRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to add new task"})
	}

	return c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid params"})
	}

	uintID, err := utils.StringToUint(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to convert ID type"})
	}

	var taskRequest models.TaskRequest

	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	updatedTask, err := h.TaskService.UpdateTask(taskRequest, uintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	taskID := c.Param("id")

	uintTaskID, err := utils.StringToUint(taskID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to convert ID type"})
	}

	if err := h.TaskService.DeleteTask(uintTaskID); err != nil {
		if err.Error() == "0" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "no task found with the given ID"})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
		}
	}

	return c.NoContent(http.StatusOK)
}
