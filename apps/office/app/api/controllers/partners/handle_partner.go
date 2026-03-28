package partners

import (
	pbPartners "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"buybikeshop/libs/go/transport"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handlePartner(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("partnerId"), 10, 64)
	result, err := s.datasource.PartnersClient.Partner(c.Request.Context(), &pbPartners.PartnerRequest{
		PartnerId: id,
	})
	if err != nil {
		log.Println(err)
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}
	transport.Success(c, result.Partner)
}
