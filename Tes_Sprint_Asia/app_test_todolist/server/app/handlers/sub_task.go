package handlers

import (
	"net/http"
	"todolist_sprint_asia/app/models"
	"todolist_sprint_asia/app/services"
	"todolist_sprint_asia/app/utils"

	"github.com/labstack/echo/v4"
)

type SubTaskHandler struct {
	SubTaskService services.SubTaskServiceInterface
}

func GetSubTaskHandler(subTaskService services.SubTaskServiceInterface) *SubTaskHandler {
	return &SubTaskHandler{SubTaskService: subTaskService}
}

func (h *SubTaskHandler) GetSubTaskByID(c echo.Context) error {
	subTaskId := c.Param("id")

	uintSubTaskId, err := utils.StringToUint(subTaskId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to convert ID type"})
	}

	task, err := h.SubTaskService.GetSubTaskByID(uintSubTaskId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get sub_task"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *SubTaskHandler) GetAllSubTaskByTaskID(c echo.Context) error {
	taskID := c.Param("taskId")

	uintTaskID, err := utils.StringToUint(taskID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to convert ID type"})
	}

	subTasks, err := h.SubTaskService.GetAllSubTaskByTaskID(uintTaskID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get all sub_task"})
	}

	return c.JSON(http.StatusOK, subTasks)
}

func (h *SubTaskHandler) AddSubTask(c echo.Context) error {
	var subTaskRequest models.SubTaskRequest

	if err := c.Bind(&subTaskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if subTaskRequest.Description == nil || *subTaskRequest.Description == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "description cannot be empty"})
	}

	createdSubTask, err := h.SubTaskService.AddSubTask(subTaskRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to add new sub task"})
	}

	return c.JSON(http.StatusCreated, createdSubTask)
}

func (h *SubTaskHandler) UpdateSubTask(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid params"})
	}

	uintID, err := utils.StringToUint(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to convert ID type"})
	}

	var subTaskRequest models.SubTaskRequest

	if err := c.Bind(&subTaskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	updatedSubTask, err := h.SubTaskService.UpdateSubTask(subTaskRequest, uintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update sub task"})
	}

	return c.JSON(http.StatusOK, updatedSubTask)
}

func (h *SubTaskHandler) DeleteSubTask(c echo.Context) error {
	subTaskId := c.Param("id")

	uintSubTaskId, err := utils.StringToUint(subTaskId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to convert ID type"})
	}

	if err := h.SubTaskService.DeleteSubTask(uintSubTaskId); err != nil {
		if err.Error() == "0" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "no sub task found with the given ID"})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete sub task"})
		}
	}

	return c.NoContent(http.StatusOK)
}
