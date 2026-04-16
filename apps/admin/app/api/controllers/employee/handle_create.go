package employee

import (
	"buybikeshop/apps/admin/app/models"
	"buybikeshop/apps/admin/app/services/employee"
	"buybikeshop/libs/go/transport"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s EmployeeController) handleCreate(ctx *gin.Context) {
	data := employee.EmployeeCreate{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		transport.Error(ctx, http.StatusBadRequest, err)
		return
	}
	data.ApplyDefaults()

	if err := s.employeeService.Validate(ctx, data); err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	empl, err := s.employeeService.Create(ctx, data)
	if err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	// TODO:
	// 	publish event instead
	go func(employee *models.Employee) {
		_, err = s.permissionService.Write(ctx, "Role", data.Role, "member", employee.Id())
		if err != nil {
			fmt.Println(err)
		}
	}(empl)

	transport.Success(ctx, empl)
}
