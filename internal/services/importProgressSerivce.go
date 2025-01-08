package services

import (
	"errors"
	"gorm.io/gorm"
	"testTask_employeeAPI/internal/models"
)

type ImportService struct {
	db *gorm.DB
}

func NewImportService(db *gorm.DB) *ImportService {
	return &ImportService{db: db}
}

func (s *ImportService) GetUncompletedImport() (*[]models.Progress, error) {
	var imp []models.Progress
	err := s.db.Where("completed = ?", false).Order("created_at desc").Find(&imp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &imp, err
}

func (s *ImportService) GetImportById(id int) (*models.Progress, error) {
	var imp models.Progress
	result := s.db.Where("id = ? AND completed = ?", id, false).First(&imp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &imp, result.Error
}

func (s *ImportService) SaveImport(imp *models.Progress) error {
	return s.db.Save(imp).Error
}
