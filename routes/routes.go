package routes

import (
	"time-tracker/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// USERS
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.GET("/users", controllers.ListUsers)

	// TASKS
	e.POST("/tasks", controllers.StartTask)
	e.PUT("/tasks/:id/end", controllers.EndTask)
	e.GET("/users/:user_id/tasks", controllers.GetUserTasks)
}
