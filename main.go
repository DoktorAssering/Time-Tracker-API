package main

import (
	"os"
	"time-tracker/config"
	"time-tracker/database"
	"time-tracker/routes"

	_ "time-tracker/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func initLogger() *logrus.Logger {
	log := logrus.New()

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

var log = initLogger()

func main() {
	config.InitConfig()
	database.ConnectDatabase()

	e := echo.New()

	// Configure logger for Echo
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","method":"${method}","uri":"${uri}","status":${status},"latency":"${latency}"}` + "\n",
		Output: log.Out,
	}))

	e.Use(middleware.Recover())

	routes.InitRoutes(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
