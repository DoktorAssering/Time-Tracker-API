package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time-tracker/config"
	"time-tracker/database"
	"time-tracker/models"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func AddUser(user *models.User) error {
	log.WithFields(logrus.Fields{
		"passportNumber": user.PassportNumber,
	}).Info("Adding user")

	// Получаем серию и номер паспорта
	passportParts := strings.Split(user.PassportNumber, " ")
	if len(passportParts) != 2 {
		return errors.New("invalid passport number format")
	}
	passportSerie := passportParts[0]
	passportNumber := passportParts[1]

	// Делаем запрос к внешнему API
	passportInfo, err := getPassportInfo(passportSerie, passportNumber)
	if err != nil {
		log.Error(err)
		return err
	}

	// Обогащаем данные пользователя
	user.Surname = passportInfo.Surname
	user.Name = passportInfo.Name
	user.Patronymic = passportInfo.Patronymic
	user.Address = passportInfo.Address

	// Сохраняем пользователя в базе данных
	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return result.Error
}

func getPassportInfo(passportSerie, passportNumber string) (*models.User, error) {
	url := fmt.Sprintf("%s/info?passportSerie=%s&passportNumber=%s", config.AppConfig.APIUrl, passportSerie, passportNumber)
	log.WithFields(logrus.Fields{
		"url": url,
	}).Info("Fetching passport info")

	resp, err := http.Get(url)
	if err != nil {
		log.Error("HTTP request failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to fetch passport info: status code %d", resp.StatusCode)
		log.Error(err)
		return nil, err
	}

	var passportInfo models.User
	err = json.NewDecoder(resp.Body).Decode(&passportInfo)
	if err != nil {
		log.Error("JSON decode failed:", err)
		return nil, err
	}

	log.WithFields(logrus.Fields{
		"surname":    passportInfo.Surname,
		"name":       passportInfo.Name,
		"patronymic": passportInfo.Patronymic,
		"address":    passportInfo.Address,
	}).Info("Fetched passport info successfully")

	return &passportInfo, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return &user, result.Error
}

func UpdateUser(user *models.User) error {
	result := database.DB.Save(&user)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return result.Error
}

func DeleteUser(id uint) error {
	result := database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return result.Error
}

func ListUsers(filters map[string]string, page, pageSize int) ([]models.User, error) {
	var users []models.User
	query := database.DB

	for key, value := range filters {
		if value != "" {
			query = query.Where(fmt.Sprintf("%s ILIKE ?", key), "%"+value+"%")
		}
	}

	offset := (page - 1) * pageSize
	result := query.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return users, result.Error
}
