package comands

import (
	"encoding/csv"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
	"strconv"
	"testTask_employeeAPI/internal/di"
	"testTask_employeeAPI/internal/models"
	"testTask_employeeAPI/pkd/Logger"
)

var db *gorm.DB
var container *di.Container

func NewCmd(c *di.Container) *cobra.Command {
	container = c

	var rootCmd = &cobra.Command{
		Use:   "import",
		Short: "Import employee data",
	}

	loadCsvCmd := &cobra.Command{
		Use:   "create-all [file]",
		Short: "Import new employees from CSV",
		Args:  cobra.ExactArgs(1),
		Run:   loadCSVData,
	}

	checkUncompletedCmd := &cobra.Command{
		Use:   "check",
		Short: "Check uncompleted import",
		Run:   checkUncompleted,
	}

	finishUncompletedCmd := &cobra.Command{
		Use:   "recover [id]",
		Short: "Check uncompleted import",
		Args:  cobra.ExactArgs(1),
		Run:   recoverImport,
	}

	rootCmd.AddCommand(loadCsvCmd)
	rootCmd.AddCommand(checkUncompletedCmd)
	rootCmd.AddCommand(finishUncompletedCmd)

	return rootCmd
}

func checkUncompleted(cmd *cobra.Command, _ []string) {
	uncompleted, err := container.ImportService.GetUncompletedImport()

	if err != nil {
		Logger.Fatal("failed to check import: %v", err)
	}

	if uncompleted == nil || len(*uncompleted) == 0 {
		Logger.Log("No uncompleted import")
		return
	}

	Logger.Log("Uncompleted import: %s", len(*uncompleted))

	for _, progress := range *uncompleted {
		Logger.Log("Uncompleted import with ID: %v. Completed: %d / %d", progress.ID, progress.Completed, progress.Total)
	}
}

func recoverImport(cmd *cobra.Command, args []string) {
	db = container.DB
	var err error

	progressId, err := strconv.Atoi(args[0])
	if err != nil {
		Logger.Fatal("failed to recover import: %v", err)
	}
	progress, err := container.ImportService.GetImportById(progressId)
	if err != nil {
		Logger.Fatal("failed to recover import: %v", err)
	}

	fileName := progress.FileName
	err = importNewEmployees(fileName, progress)
	if err != nil {
		Logger.Fatal("failed to recover import: %v", err)
	}

	Logger.Log("Successfully imported CSV data")
}

func loadCSVData(cmd *cobra.Command, args []string) {
	db = container.DB
	var err error

	// Загрузка данных из CSV
	fileName := args[0]
	err = importNewEmployees(fileName, nil)
	if err != nil {
		Logger.Fatal("failed to load CSV data: %v", err)
	}

	Logger.Log("Successfully imported CSV data")
}

func importNewEmployees(fileName string, progress *models.Progress) error {
	file, err := os.Open(fileName)
	if err != nil {
		return Logger.Error("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return Logger.Error("failed to read CSV file: %w", err)
	}

	// Пропуск заголовков
	records = records[1:]
	if progress == nil {
		progress = &models.Progress{
			FileName: fileName,
			Total:    len(records) - 1,
		}
	} else {
		// Пропуск уже импортирвоанных сотрудников
		records = records[progress.ProcessedCount:]
	}

	// Обработка каждой строки в CSV
	length := len(records)
	for index, record := range records {
		employee, err := loadRow(record)
		if err != nil {
			return Logger.Error("failed to insert employee: %w", err)
		}

		progress.ProcessedCount++
		container.ImportService.SaveImport(progress)

		logProgress(index, employee, length)
	}

	progress.Completed = true
	err = container.ImportService.SaveImport(progress)
	if err != nil {
		return Logger.Error("failed to save CSV file: failed to save import progress: %w", err)
	}

	return nil
}

func logProgress(index int, employee *models.Employee, length int) {
	Logger.Log("[Employees import] Progress %d / %d. Employee with Id %d and name '%s' successfully imported",
		index, length, employee.Id, employee.Name)
}

func loadRow(record []string) (*models.Employee, error) {
	name := record[0]
	jobTitle := record[1]
	department := record[2]
	partTime := false

	// P - partTime, F - fullTime
	if record[3] == models.PartTimeValue {
		partTime = true
	} else if len(record[3]) != 0 && record[3] != models.FullTimeValue {
		return nil, Logger.Error("invalid value '%s' in column 'Full or Part-Time', expected 'P' or 'F''", record[3])
	}

	// Создание сотрудника
	employee := &models.Employee{
		Name:     name,
		PartTime: partTime,
	}

	jobId, err := container.JobService.SaveOrCreateJob(jobTitle)
	if err != nil {
		return nil, err
	}
	employee.ID = jobId

	departmentId, err := container.DepartmentService.SaveOrCreateDepartment(department)
	if err != nil {
		return nil, err
	}

	employee.ID = departmentId

	// Добавляем сотрудника в базу
	err = db.Create(&employee).Error
	if err != nil {
		return nil, Logger.Error("failed to insert employee: %w", err)
	}

	salaryOrHourly := record[4]
	typicalHours, _ := strconv.ParseFloat(record[5], 64)
	annualSalary, _ := strconv.ParseFloat(record[6], 64)
	hourlyRate, _ := strconv.ParseFloat(record[7], 64)
	err = container.IncomeService.SaveOrCreateIncome(salaryOrHourly, typicalHours, annualSalary, hourlyRate, employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
