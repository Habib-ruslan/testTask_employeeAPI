package services

import (
	"errors"
	"gorm.io/gorm"
	"testTask_employeeAPI/internal/models"
	"testTask_employeeAPI/pkd/Logger"
)

type DepartmentService struct {
	Db *gorm.DB
}

func NewDepartmentService(db *gorm.DB) *DepartmentService {
	return &DepartmentService{Db: db}
}

func (service *DepartmentService) SaveOrCreateDepartment(department string) (uint, error) {
	var dept models.Department
	err := service.Db.Where("name = ?", department).First(&dept).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, Logger.Error("failed to find department: %w", err)
	}
	if dept.Id == 0 {
		dept = models.Department{Name: department}
		service.Db.Create(&dept)
	}

	return dept.Id, nil
}
