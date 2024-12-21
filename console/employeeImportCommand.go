package loadcsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"testTask_employeeAPI/app"
	"testTask_employeeAPI/models"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "loadcsv [file]",
		Short: "Load employee data from CSV",
		Args:  cobra.ExactArgs(2),
		Run:   loadCSVData,
	}
}

func loadCSVData(cmd *cobra.Command, args []string) {
	db = app.GetApp().DB
	var err error

	// Загрузка данных из CSV
	fileName := args[1]
	err = loadCSV(fileName)
	if err != nil {
		log.Fatalf("failed to load CSV data: %v", err)
	}

	fmt.Println("Data loaded successfully!")
}

func loadCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	// Пропуск заголовков
	records = records[1:]

	// Обработка каждой строки в CSV
	for _, record := range records {
		err = loadRow(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadRow(record []string) error {
	name := record[0]
	jobTitle := record[1]
	department := record[2]
	partTime := false

	// P - partTime, F - fullTime
	if record[3] == models.PartTimeValue {
		partTime = true
	} else if len(record[3]) != 0 && record[3] != models.FullTimeValue {
		return fmt.Errorf("invalid value '%s' in column 'Full or Part-Time', expected 'P' or 'F''", record[3])
	}

	// Создание сотрудника
	employee := &models.Employee{
		Name:     name,
		PartTime: partTime,
	}

	err := SaveOrCreateJob(jobTitle, employee)
	if err != nil {
		return err
	}

	err = SaveOrCreateDepartment(department, employee)
	if err != nil {
		return err
	}

	// Добавляем сотрудника в базу
	err = db.Create(&employee).Error
	if err != nil {
		return fmt.Errorf("failed to insert employee: %w", err)
	}

	err = SaveOrCreateIncome(record, employee)
	if err != nil {
		return err
	}

	return nil
}

func SaveOrCreateJob(jobTitle string, employee *models.Employee) error {
	var job models.Job
	err := db.Where("name = ?", jobTitle).First(&job).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find job: %w", err)
	}
	if job.Id == 0 {
		job = models.Job{Name: jobTitle}
		db.Create(&job)
	}

	employee.JobId = job.Id

	return nil
}

func SaveOrCreateDepartment(department string, employee *models.Employee) error {
	var dept models.Department
	err := db.Where("name = ?", department).First(&dept).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find department: %w", err)
	}
	if dept.Id == 0 {
		dept = models.Department{Name: department}
		db.Create(&dept)
	}
	employee.DepartmentId = dept.Id

	return nil
}

func SaveOrCreateIncome(record []string, employee *models.Employee) error {
	salaryOrHourly := record[4]
	typicalHours, _ := strconv.ParseFloat(record[5], 64)
	annualSalary, _ := strconv.ParseFloat(record[6], 64)
	hourlyRate, _ := strconv.ParseFloat(record[7], 64)

	var salary *models.Salary
	var hourly *models.Hourly
	var err error = nil
	if salaryOrHourly == models.SalaryMode || int(typicalHours) == 0 {
		salary = &models.Salary{
			AnnualSalary: annualSalary,
			EmployeeId:   employee.Id,
		}
		err = db.Create(&salary).Error
	} else if salaryOrHourly == models.HourlyMode {
		hourly = &models.Hourly{
			HourlyRate:   float32(hourlyRate),
			TypicalHours: int(typicalHours),
			EmployeeId:   employee.Id,
		}
		err = db.Create(&hourly).Error
	} else {
		return fmt.Errorf("invalid value '%s' in column 'SalaryOrHourly'", record[4])
	}

	return err
}
