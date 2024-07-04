package controllers

import (
	"net/http"
	"strconv"
	"time"
	"time-tracker/models"
	"time-tracker/services"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var taskLog = logrus.New()

func StartTask(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		taskLog.Error("Failed to bind task:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	newTask, err := services.StartTask(task.UserID, task.TaskName)
	if err != nil {
		taskLog.Error("Failed to start task:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, newTask)
}

func EndTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := services.EndTask(uint(id))
	if err != nil {
		taskLog.Warn("Task not found:", err)
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

func GetUserTasks(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	start, err := time.Parse(time.RFC3339, c.QueryParam("start"))
	if err != nil {
		taskLog.Error("Invalid start date:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid start date"})
	}
	end, err := time.Parse(time.RFC3339, c.QueryParam("end"))
	if err != nil {
		taskLog.Error("Invalid end date:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid end date"})
	}

	tasks, err := services.GetUserTasks(uint(userID), start, end)
	if err != nil {
		taskLog.Error("Failed to get user tasks:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, tasks)
}
