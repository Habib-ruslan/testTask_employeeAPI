package main

import (
	"fmt"
	"log"
	"testTask_employeeAPI/configs"
	"testTask_employeeAPI/database"
	loadcsv "testTask_employeeAPI/internal/cli/comands"
	"testTask_employeeAPI/internal/di"
	"testTask_employeeAPI/internal/models"
)

func main() {
	configure()
	runConsole() // Консольная команда
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
		&models.Progress{},
	); err != nil {
		log.Fatalf("Error migrating employee table: %v", err)
	}

	di.NewContainer(db, *conf)
}

func runConsole() {
	cmd := loadcsv.NewCmd(di.Current)
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}
