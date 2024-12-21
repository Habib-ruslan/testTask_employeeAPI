package routes

import (
	"github.com/gin-gonic/gin"
	"testTask_employeeAPI/app"
	"testTask_employeeAPI/controllers"
	"testTask_employeeAPI/services"
)

func RegisterRoutes(router *gin.Engine) {
	employeeService := services.EmployeeService{DB: app.GetApp().DB}
	employee := router.Group("/employee")
	{
		employee.GET("/:name", controllers.NewEmployeeController(employeeService).GetEmployee)
	}
}
