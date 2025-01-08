package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testTask_employeeAPI/internal/models"
	"testTask_employeeAPI/internal/services"
)

const DefaultPageSize = 10
const DefaultPage = 1

type EmployeeController struct {
	EmployeeService services.EmployeeService
}

type RequestParams struct {
	Search   string `json:"search" binding:"max=255"`
	PageSize int    `json:"page_size" binding:"min=1"`
	Page     int    `json:"page" binding:"min=1"`
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

func (controller *EmployeeController) GetEmployee(c *gin.Context) {
	params := RequestParams{
		Page:     DefaultPage,
		PageSize: DefaultPageSize,
	}
	params.Search = c.DefaultQuery("search", "")
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(params.PageSize)))
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size"})
		return
	}
	params.PageSize = pageSize

	page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(params.Page)))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	params.Page = page

	var employees *[]models.Employee
	employees, err = controller.EmployeeService.GetAllEmployeesBySearch(params.Search, params.PageSize, params.Page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var result []*EmployeeResponse
	for _, employee := range *employees {
		result = append(result, ToEmployeeResponse(&employee))
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
