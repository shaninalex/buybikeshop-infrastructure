package partners

import (
	pbPartners "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handleRolesList(c *gin.Context) {
	result, err := s.datasource.PartnersClient.PartnerRoleList(c.Request.Context(), &pbPartners.PartnerRoleListRequest{
		Query: c.Request.URL.RawQuery, // NOTE: grpc should now know about query params, it has to get only parsed list of search params
	})
	if err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}
	transport.Success(c, result.Roles)
}
