package services

import (
	"time"
	"time-tracker/database"
	"time-tracker/models"

	"github.com/sirupsen/logrus"
)

var taskLog = logrus.New()

func StartTask(userID uint, taskName string) (*models.Task, error) {
	task := &models.Task{
		UserID:    userID,
		TaskName:  taskName,
		StartTime: time.Now(),
	}
	result := database.DB.Create(&task)
	if result.Error != nil {
		taskLog.Error(result.Error)
	}
	return task, result.Error
}

func EndTask(taskID uint) (*models.Task, error) {
	var task models.Task
	if err := database.DB.First(&task, taskID).Error; err != nil {
		taskLog.Error(err)
		return nil, err
	}

	endTime := time.Now()
	task.EndTime = &endTime
	result := database.DB.Save(&task)
	if result.Error != nil {
		taskLog.Error(result.Error)
	}
	return &task, result.Error
}

func GetUserTasks(userID uint, start, end time.Time) ([]models.Task, error) {
	var tasks []models.Task
	result := database.DB.Where("user_id = ? AND start_time >= ? AND start_time <= ?", userID, start, end).Order("end_time - start_time desc").Find(&tasks)
	if result.Error != nil {
		taskLog.Error(result.Error)
	}
	return tasks, result.Error
}
