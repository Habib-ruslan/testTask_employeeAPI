package services

import (
	"gorm.io/gorm"
	"testTask_employeeAPI/models"
)

type EmployeeService struct {
	DB *gorm.DB
}

func (service *EmployeeService) GetAllEmployeesByName(name string) (*[]models.Employee, error) {
	var employees []models.Employee
	if err := service.DB.Where("employees.name = ?", name).
		Preload("Job").Preload("Department").Preload("Salary").Preload("Hourly").
		Find(&employees).Error; err != nil {
		return nil, err
	}

	return &employees, nil
}
