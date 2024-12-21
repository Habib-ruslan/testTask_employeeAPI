package app

import (
	"gorm.io/gorm"
	"testTask_employeeAPI/configs"
)

var app *Application

type Application struct {
	Config *configs.Config
	DB     *gorm.DB
}

func Init(config *configs.Config, db *gorm.DB) {
	app = &Application{config, db}
}

func GetApp() *Application {
	return app
}
