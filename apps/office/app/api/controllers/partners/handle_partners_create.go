package partners

import (
	pbPartners "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"buybikeshop/libs/go/transport"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *Controller) handlePartnersCreate(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		transport.Error(c, http.StatusBadRequest, err)
		return
	}

	var partner pbPartners.Partner
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(body, &partner); err != nil {
		transport.Error(c, http.StatusBadRequest, err)
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
