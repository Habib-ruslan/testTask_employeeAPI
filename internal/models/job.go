package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255" json:"name"`
}
