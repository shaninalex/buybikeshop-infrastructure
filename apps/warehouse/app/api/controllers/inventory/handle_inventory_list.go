package inventory

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *InventoryController) handleInventoryList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
