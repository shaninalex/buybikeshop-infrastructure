package partners

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
	router.GET("partners", s.handlePartnersList)
	router.GET("partners/:partnerId", s.handlePartner)
	router.POST("partners", s.handlePartnersCreate)
	router.GET("roles", s.handleRolesList)
}
