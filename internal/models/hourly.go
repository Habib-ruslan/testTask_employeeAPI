package models

import "gorm.io/gorm"

type Hourly struct {
	gorm.Model
	Id           uint    `gorm:"primary_key" json:"id"`
	EmployeeId   uint    `json:"employee_id"`
	TypicalHours int     `gorm:"default:0" json:"typical_hours"`
	HourlyRate   float32 `gorm:"type:numeric(10,2);default:0" json:"hourly_rate"`
}
