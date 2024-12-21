package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255" json:"name"`
}
