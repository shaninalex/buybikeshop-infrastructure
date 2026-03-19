package product

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"buybikeshop/libs/go/transport"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrorEmptyProductId = errors.New("empty product id")
)

func (s *ProductController) handleProductVariants(c *gin.Context) {
	strId := c.Param("product_id")
	if strId == "" {
		transport.Error(c, http.StatusInternalServerError, ErrorEmptyProductId)
		return
	}
	productId, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}

	data, err := s.datasource.CatalogClient.ProductVariantList(c.Request.Context(), &pb.ProductVariantListRequest{
		ProductIds: []uint64{productId},
	})
	if err != nil {
		transport.Error(c, http.StatusInternalServerError, err)
		return
	}

	transport.Success(c, data.GetVariants())
}
