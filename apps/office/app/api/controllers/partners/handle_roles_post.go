package partners

import (
	pbPartners "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handleRolesPost(c *gin.Context) {
	var payload pbPartners.PartnerRole
	if err := c.ShouldBindJSON(&payload); err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}

	result, err := s.datasource.PartnersClient.PartnerRoleSave(c.Request.Context(), &pbPartners.PartnerRoleSaveRequest{
		Role: &payload,
	})
	if err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}
	transport.Success(c, result.Role)
}
