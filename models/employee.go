package models

import "gorm.io/gorm"

const FullTimeValue string = "F"
const PartTimeValue string = "P"
const SalaryMode string = "SALARY"
const HourlyMode string = "HOURLY"

type Employee struct {
	gorm.Model
	Id           uint       `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"index;size:255" json:"name"`
	JobId        uint       `json:"job_id"`
	Job          Job        `gorm:"foreignKey:JobId" json:"job"`
	DepartmentId uint       `json:"department_id"`
	Department   Department `gorm:"DepartmentId:JobId" json:"department"`
	Salary       *Salary    `gorm:"foreignKey:EmployeeId" json:"salary"`
	Hourly       *Hourly    `gorm:"foreignKey:EmployeeId" json:"hourly"`
	PartTime     bool       `gorm:"part_time" json:"part_time"`
}
