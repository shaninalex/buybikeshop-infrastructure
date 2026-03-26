package api

import (
	"buybikeshop/apps/office"
	"buybikeshop/apps/office/app/api/controllers/catalog"
	"buybikeshop/apps/office/app/api/controllers/partners"
	"buybikeshop/libs/go/auth"
	"buybikeshop/libs/go/config"
	"net/http"

	ory "github.com/ory/kratos-client-go"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func ProvideRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.DebugMode)

	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.GET("/_health", HealthRoute)

	return router
}

func HealthRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"app":     "office",
		"version": office.Version,
	})
}

type ApiDeps struct {
	dig.In

	KratosAPIClient    *ory.APIClient
	Config             *config.Config
	CatalogController  *catalog.Controller
	PartnersController *partners.Controller
}

func ProvideAPI(deps ApiDeps) *gin.Engine {
	router := ProvideRouter()
	router.Use(gin.Recovery())

	router.Use(auth.AuthMiddleware(deps.KratosAPIClient, deps.Config))

	v1 := router.Group("/api/v1/office")
	deps.CatalogController.Register(v1)
	deps.PartnersController.Register(v1)

	return router
}
