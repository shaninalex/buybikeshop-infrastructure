package product

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"

	"github.com/gin-gonic/gin"
)

func (s *ProductController) handleProducts(c *gin.Context) {
	data, err := s.datasource.CatalogClient.ProductList(c.Request.Context(), &pb.ProductListRequest{})
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, data.Products)
}
