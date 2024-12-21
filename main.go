package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"testTask_employeeAPI/app"
	"testTask_employeeAPI/configs"
	loadcsv "testTask_employeeAPI/console"
	"testTask_employeeAPI/database"
	"testTask_employeeAPI/models"
	"testTask_employeeAPI/routes"
)

func main() {
	configure()

	if len(os.Args) > 1 && os.Args[1] == "loadcsv" {
		runConsole() // Консольная команда
	} else {
		runWeb() // Веб-приложение
	}
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

	app.Init(conf, db)
}

func runWeb() {
	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(configs.GetConfig().Server.Port)
}

func runConsole() {
	cmd := loadcsv.NewCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}
