package employee

import (
	"buybikeshop/apps/admin/app/services/employee"
	"buybikeshop/libs/go/transport"
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

	// TODO: get from payload, not hardcode!
	data.Group = "manager"
	data.Department = "office"

	if err := s.employeeService.Validate(ctx, data); err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	empl, err := s.employeeService.Create(ctx, data)
	if err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	transport.Success(ctx, empl)
}
