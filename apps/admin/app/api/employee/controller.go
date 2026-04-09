package employee

import "github.com/gin-gonic/gin"

type EmployeeController struct {
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{}
}

func (s *EmployeeController) Register(router *gin.RouterGroup) {
	router.GET("inventory", nil)
}
