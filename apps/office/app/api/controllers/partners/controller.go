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
	router.GET("partners/roles", s.handleRolesList)
	router.POST("partners/roles", s.handleRolesPost)
	router.PATCH("partners/roles/:roleId", s.handleRolesPatch)
	router.GET("partners/:partnerId", s.handlePartner)
	router.POST("partners", s.handlePartnersCreate)
}
