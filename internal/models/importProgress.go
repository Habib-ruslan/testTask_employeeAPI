package models

import "gorm.io/gorm"

type Progress struct {
	gorm.Model
	Id             uint   `gorm:"primaryKey"`
	FileName       string `json:"file_name"`
	ProcessedCount int    `json:"processed_count"`
	Total          int    `json:"total"`
	Completed      bool   `json:"completed"`
}
