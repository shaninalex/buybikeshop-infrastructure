package employee

import (
	"buybikeshop/apps/admin/app/services/employee"
	"buybikeshop/libs/go/keto"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService   employee.Service
	permissionService keto.KetoWriter
}

func NewEmployeeController(
	employeeService employee.Service,
	permissionService keto.KetoWriter,
) *EmployeeController {
	return &EmployeeController{
		employeeService:   employeeService,
		permissionService: permissionService,
	}
}

func (s *EmployeeController) Register(router *gin.RouterGroup) {
	router.GET("employees", s.handleList)
	router.POST("employees/create", s.handleCreate)
	router.PATCH("employees/:id", s.handlePatch)
}
