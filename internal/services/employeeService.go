package services

import (
	"gorm.io/gorm"
	"strings"
	"testTask_employeeAPI/internal/models"
)

type EmployeeService struct {
	DB *gorm.DB
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	return &EmployeeService{
		DB: db,
	}
}

func (service *EmployeeService) GetAllEmployeesBySearch(search string, limit int, offset int) (*[]models.Employee, error) {
	var employees []models.Employee

	searchTerms := strings.Fields(search)

	query := service.DB
	for _, term := range searchTerms {
		likePattern := "%" + term + "%"
		query = query.Where("employees.name LIKE ?", likePattern)
	}

	// Пагинация
	query = query.Limit(limit).Offset((offset - 1) * limit)

	if err := query.Preload("Job").
		Preload("Department").
		Preload("Salary").
		Preload("Hourly").
		Find(&employees).Error; err != nil {
		return nil, err
	}

	return &employees, nil
}
