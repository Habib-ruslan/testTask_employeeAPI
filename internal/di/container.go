package di

import (
	"gorm.io/gorm"
	"testTask_employeeAPI/configs"
	"testTask_employeeAPI/internal/http/handlers"
	"testTask_employeeAPI/internal/services"
)

var Current *Container

type Container struct {
	DB                 *gorm.DB
	Config             *configs.Config
	EmployeeService    *services.EmployeeService
	EmployeeController *handlers.EmployeeController
	IncomeService      *services.IncomeService
	DepartmentService  *services.DepartmentService
	JobService         *services.JobService
	ImportService      *services.ImportService
}

func NewContainer(db *gorm.DB, config configs.Config) *Container {
	employeeService := services.NewEmployeeService(db)
	employeeHandler := handlers.NewEmployeeController(*employeeService)

	Current := &Container{
		DB:                 db,
		Config:             &config,
		EmployeeService:    employeeService,
		EmployeeController: employeeHandler,
		IncomeService:      services.NewIncomeService(db),
		DepartmentService:  services.NewDepartmentService(db),
		JobService:         services.NewJobService(db),
		ImportService:      services.NewImportService(db),
	}

	return Current
}
