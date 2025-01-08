package main

import (
	"log"
	"testTask_employeeAPI/configs"
	"testTask_employeeAPI/database"
	"testTask_employeeAPI/internal/di"
	"testTask_employeeAPI/internal/http/routes"
	"testTask_employeeAPI/internal/models"
)

func main() {
	configure()
}

func configure() {
	configs.LoadConfig()
	conf := configs.GetConfig()

	db := database.InitConnection(conf.Database)

	if err := db.AutoMigrate(
		&models.Employee{},
		&models.Department{},
		&models.Salary{},
		&models.Hourly{},
		&models.Job{},
	); err != nil {
		log.Fatalf("Error migrating employee table: %v", err)
	}

	container := di.NewContainer(db, *conf)

	r := routes.RegisterRoutes(container)
	r.Run(configs.GetConfig().Server.Port)
}
