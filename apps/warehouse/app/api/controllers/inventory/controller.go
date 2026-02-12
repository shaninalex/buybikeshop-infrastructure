package inventory

import "github.com/gin-gonic/gin"

type InventoryController struct {
}

func NewInventoryController() *InventoryController {
	return &InventoryController{}
}

func (s *InventoryController) Register(router *gin.RouterGroup) {
	router.GET("inventory", s.handleInventoryList)
}
