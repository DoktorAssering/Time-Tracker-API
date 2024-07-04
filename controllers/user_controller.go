package controllers

import (
	"net/http"
	"strconv"
	"time-tracker/models"
	"time-tracker/services"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var log = logrus.New()

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		log.Error("Failed to bind user:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	err := services.AddUser(user)
	if err != nil {
		log.Error("Failed to add user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn("User not found:", err)
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}
		log.Error("Failed to get user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn("User not found:", err)
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}
		log.Error("Failed to get user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if err := c.Bind(user); err != nil {
		log.Error("Failed to bind user:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	err = services.UpdateUser(user)
	if err != nil {
		log.Error("Failed to update user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteUser(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn("User not found:", err)
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}
		log.Error("Failed to delete user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func ListUsers(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
		pageSize = 10
	}

	filters := make(map[string]string)
	for _, param := range []string{"PassportNumber", "Surname", "Name", "Patronymic", "Address"} {
		if value := c.QueryParam(param); value != "" {
			filters[param] = value
		}
	}

	users, err := services.ListUsers(filters, page, pageSize)
	if err != nil {
		log.Error("Failed to list users:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
