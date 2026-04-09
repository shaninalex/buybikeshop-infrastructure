package employee

import (
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s EmployeeController) handleList(ctx *gin.Context) {
	employees, err := s.employeeService.List(ctx)
	if err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	transport.Success(ctx, employees)
}
