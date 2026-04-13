package employee

import (
	"buybikeshop/apps/admin/app/models"
	"buybikeshop/libs/go/kratos"
	"buybikeshop/libs/go/transport"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s EmployeeController) handleCreate(ctx *gin.Context) {
	// TODO: kratos only responsible for identities! Not employers
	data := kratos.EmployeeCreate{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		transport.Error(ctx, http.StatusBadRequest, err)
		return
	}
	data.ApplyDefaults()
	// TODO:
	// 	validate method should check uniqueness email, uniqueness phone, proper phone format, future DOB etc.
	// 	s.employeeService.Validate(ctx, data)
	employee, err := s.employeeService.Create(ctx, data)
	if err != nil {
		transport.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	// TODO:
	// 	publish event instead
	go func(employee *models.Employee) {
		// TODO: values should be from payload
		_, err = s.permissionService.Write(ctx, "Role", "manager", "member", employee.Id())
		if err != nil {
			fmt.Println(err)
		}
	}(employee)

	transport.Success(ctx, employee)
}
