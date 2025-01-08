package services

import (
	"gorm.io/gorm"
	"testTask_employeeAPI/internal/models"
	"testTask_employeeAPI/pkd/Logger"
)

type IncomeService struct {
	Db *gorm.DB
}

func NewIncomeService(db *gorm.DB) *IncomeService {
	return &IncomeService{Db: db}
}

func (service *IncomeService) SaveOrCreateIncome(salaryOrHourly string, typicalHours, annualSalary, hourlyRate float64, employee *models.Employee) error {
	var err error = nil
	if salaryOrHourly == models.SalaryMode || int(typicalHours) == 0 {
		err = service.saveOrCreateSalary(annualSalary, employee)
		if err != nil {
			return err
		}
	} else if salaryOrHourly == models.HourlyMode {
		err = service.saveOrCreateHourly(typicalHours, hourlyRate, employee)
		if err != nil {
			return err
		}
	} else {
		return Logger.Error("invalid value '%s' in column 'SalaryOrHourly'", salaryOrHourly)
	}

	return err
}

func (service *IncomeService) saveOrCreateSalary(annualSalary float64, employee *models.Employee) error {
	salary := &models.Salary{
		AnnualSalary: annualSalary,
		EmployeeId:   employee.Id,
	}

	return service.Db.Create(&salary).Error
}

func (service *IncomeService) saveOrCreateHourly(typicalHours, hourlyRate float64, employee *models.Employee) error {
	hourly := &models.Hourly{
		HourlyRate:   float32(hourlyRate),
		TypicalHours: int(typicalHours),
		EmployeeId:   employee.Id,
	}

	return service.Db.Create(&hourly).Error
}
