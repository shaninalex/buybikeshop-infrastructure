package catalog

import (
	"buybikeshop/apps/office/app/pkg/connector"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	datasource *connector.DatasourceClient
}

func NewController(datasource *connector.DatasourceClient) *Controller {
	return &Controller{
		datasource: datasource,
	}
}

func (s *Controller) Register(router *gin.RouterGroup) {
	router.GET("products", s.handleProducts)
	router.GET("products/:product_id/variants", s.handleProductVariants)
}
