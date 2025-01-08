package services

import (
	"errors"
	"gorm.io/gorm"
	"testTask_employeeAPI/internal/models"
	"testTask_employeeAPI/pkd/Logger"
)

type JobService struct {
	Db *gorm.DB
}

func NewJobService(db *gorm.DB) *JobService {
	return &JobService{Db: db}
}

func (service *JobService) SaveOrCreateJob(jobTitle string) (uint, error) {
	var job models.Job
	err := service.Db.Where("name = ?", jobTitle).First(&job).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, Logger.Error("failed to find job: %w", err)
	}
	if job.Id == 0 {
		job = models.Job{Name: jobTitle}
		service.Db.Create(&job)
	}

	return job.Id, nil
}
