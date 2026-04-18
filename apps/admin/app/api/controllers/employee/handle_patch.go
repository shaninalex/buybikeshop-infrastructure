package employee

import (
	"buybikeshop/apps/admin/app/services/employee"
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s EmployeeController) handlePatch(ctx *gin.Context) {
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

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		transport.Error(ctx, http.StatusBadRequest, err)
		return
	}

	empl, err := s.employeeService.Patch(ctx, id, data)
	if err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	transport.Success(ctx, empl)
}
