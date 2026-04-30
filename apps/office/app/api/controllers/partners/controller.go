package partners

import (
	"buybikeshop/apps/office/app/pkg/connector"
	"buybikeshop/libs/go/keto"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	datasource        *connector.DatasourceClient
	permissionService *keto.Manager
}

func NewController(permissionService *keto.Manager, datasource *connector.DatasourceClient) *Controller {
	return &Controller{
		permissionService: permissionService,
		datasource:        datasource,
	}
}

func (s *Controller) GetNamespace() string {
	return "Partner"
}

func (s *Controller) Register(router *gin.RouterGroup) {
	router.GET("partners", s.check("read"), s.handlePartnersList)
	router.GET("partners/roles", s.handleRolesList)
	router.POST("partners/roles", s.handleRolesPost)
	router.PATCH("partners/roles/:roleId", s.handleRolesPatch)
	router.GET("partners/:partnerId", s.check("read"), s.handlePartner)
	router.POST("partners", s.check("create"), s.handlePartnersCreate)
}

func (s *Controller) check(action string) gin.HandlerFunc {
	return keto.PermissionMiddleware(s.permissionService, s.GetNamespace(), "*", action)
}
