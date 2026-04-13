package employee

import (
	"buybikeshop/libs/go/kratos"
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s EmployeeController) handleCreate(ctx *gin.Context) {
	data := kratos.EmployeeCreate{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		transport.Error(ctx, http.StatusBadRequest, err)
		return
	}
	data.ApplyDefaults()
	employees, err := s.employeeService.Create(ctx, data)
	if err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	transport.Success(ctx, employees)
}
