package handlers

import "testTask_employeeAPI/internal/models"

type EmployeeResponse struct {
	Id             uint    `json:"id"`
	Name           string  `json:"name"`
	JobTitle       string  `json:"job_title"`
	Department     string  `json:"department"`
	PartTime       string  `json:"full_or_part_time"`
	SalaryOrHourly string  `json:"salary_or_hourly"`
	TypicalHours   int     `json:"typical_hourly"`
	AnnualSalary   float64 `json:"annual_salary"`
	HourlyRate     float32 `json:"hourly_rate"`
}

func ToEmployeeResponse(employee *models.Employee) *EmployeeResponse {
	var typicalSalary int
	var hourlyRate float32
	var annualSalary float64
	salaryOrHourly := models.SalaryMode

	if employee.Hourly != nil {
		typicalSalary = employee.Hourly.TypicalHours
		hourlyRate = employee.Hourly.HourlyRate
		salaryOrHourly = models.HourlyMode
	}

	if employee.Salary != nil {
		annualSalary = employee.Salary.AnnualSalary
	}

	partTime := models.FullTimeValue
	if employee.PartTime {
		partTime = models.PartTimeValue
	}

	return &EmployeeResponse{
		Id:             employee.Id,
		Name:           employee.Name,
		JobTitle:       employee.Job.Name,
		Department:     employee.Department.Name,
		PartTime:       partTime,
		SalaryOrHourly: salaryOrHourly,
		TypicalHours:   typicalSalary,
		AnnualSalary:   annualSalary,
		HourlyRate:     hourlyRate,
	}
}
