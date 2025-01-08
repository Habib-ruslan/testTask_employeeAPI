package routes

import (
	"github.com/gin-gonic/gin"
	"testTask_employeeAPI/internal/di"
)

func RegisterRoutes(container *di.Container) *gin.Engine {
	router := gin.Default()
	employee := router.Group("/employee")
	{
		employee.GET("/", container.EmployeeController.GetEmployee)
	}

	return router
}
