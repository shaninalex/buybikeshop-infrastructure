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
	//go func(employee *models.Employee) {
	id := empl.Id()
	if err = s.permissionService.Assign(ctx, "manager", &id, nil); err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	//}(empl)

	transport.Success(ctx, empl)
}
