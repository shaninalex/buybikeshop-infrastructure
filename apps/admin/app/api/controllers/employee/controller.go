package employee

import (
	"buybikeshop/apps/admin/app/services/employee"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService employee.Service
}

func NewEmployeeController(
	employeeService employee.Service,
) *EmployeeController {
	return &EmployeeController{
		employeeService: employeeService,
	}
}

func (s *EmployeeController) Register(router *gin.RouterGroup) {
	router.GET("employees", s.handleList)
	router.POST("employees/create", s.handleCreate)
	router.PATCH("employees/:id", s.handlePatch)
}
