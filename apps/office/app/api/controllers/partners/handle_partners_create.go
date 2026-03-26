package partners

import (
	pbPartners "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handlePartnersCreate(c *gin.Context) {
	var partner pbPartners.Partner
	if err := c.BindJSON(&partner); err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}
	result, err := s.datasource.PartnersClient.PartnersSave(c.Request.Context(), &pbPartners.PartnersSaveRequest{
		Partner: &partner,
	})
	if err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}

	transport.Success(c, result.Partner)
}
