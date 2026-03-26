package partners

import (
	pbPartners "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handlePartnersList(c *gin.Context) {
	result, err := s.datasource.PartnersClient.PartnersList(c.Request.Context(), &pbPartners.PartnersListRequest{
		Query: c.Request.URL.RawQuery,
	})
	if err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}
	transport.Success(c, result.Partners)
}
