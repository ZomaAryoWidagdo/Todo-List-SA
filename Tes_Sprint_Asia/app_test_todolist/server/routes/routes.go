package rest

import (
	"net/http"
	"todolist_sprint_asia/app/handlers"
	"todolist_sprint_asia/app/repositories"
	"todolist_sprint_asia/app/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}

func InitRoute(e *echo.Echo, db *gorm.DB) {
	taskRepo := repositories.GetTaskRepository(db)
	subTaskRepo := repositories.GetSubTaskRepository(db)

	taskService := services.GetTaskService(taskRepo)
	subTaskService := services.GetSubTaskService(subTaskRepo)

	taskHandler := handlers.GetTaskHandler(taskService)
	subTaskHandler := handlers.GetSubTaskHandler(subTaskService)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodGet},
	}))

	e.GET("/ping", ping)

	e.GET("/task", taskHandler.GetAllTask)
	e.GET("/task/:id", taskHandler.GetTaskByID)
	e.POST("/task", taskHandler.AddTask)
	e.PATCH("/task/:id", taskHandler.UpdateTask)
	e.DELETE("/task/:id", taskHandler.DeleteTask)

	e.GET("/subtask/all/:taskId", subTaskHandler.GetAllSubTaskByTaskID)
	e.GET("/subtask/:id", subTaskHandler.GetSubTaskByID)
	e.POST("/subtask", subTaskHandler.AddSubTask)
	e.PATCH("/subtask/:id", subTaskHandler.UpdateSubTask)
	e.DELETE("/subtask/:id", subTaskHandler.DeleteSubTask)

}
