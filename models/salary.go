package models

type Salary struct {
	Id           uint    `gorm:"primary_key" json:"id"`
	EmployeeId   uint    `json:"employee_id"`
	AnnualSalary float64 `gorm:"default:0" json:"annual_salary"`
}
