package main

import (
	"todolist_sprint_asia/config"
	setup "todolist_sprint_asia/db"
	rest "todolist_sprint_asia/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	e := echo.New()

	db := setup.InitDB()

	rest.InitRoute(e, db)

	e.Start(":8080")
}
