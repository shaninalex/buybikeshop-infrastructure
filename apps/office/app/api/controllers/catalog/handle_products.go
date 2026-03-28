package catalog

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"buybikeshop/libs/go/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handleProducts(c *gin.Context) {
	data, err := s.datasource.CatalogClient.ProductList(c.Request.Context(), &pb.ProductListRequest{})
	if err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}

	transport.Success(c, data)
}
