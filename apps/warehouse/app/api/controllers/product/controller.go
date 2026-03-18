package product

import (
	"buybikeshop/apps/warehouse/app/pkg/connector"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	datasource *connector.DatasourceClient
}

func NewProductController(datasource *connector.DatasourceClient) *ProductController {
	return &ProductController{
		datasource: datasource,
	}
}

func (s *ProductController) Register(router *gin.RouterGroup) {
	router.GET("products", s.handleProducts)
	router.GET("products/:product_id/variants", s.handleProductVariants)
}
