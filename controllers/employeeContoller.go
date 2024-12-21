package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTask_employeeAPI/services"
)

type EmployeeController struct {
	EmployeeService services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

func (controller *EmployeeController) GetEmployee(c *gin.Context) {
	name := c.Param("name")
	employees, err := controller.EmployeeService.GetAllEmployeesByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
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
