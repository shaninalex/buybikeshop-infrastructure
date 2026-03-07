package api

import (
	"buybikeshop/apps/warehouse"
	"buybikeshop/apps/warehouse/app/api/controllers/inventory"
	"buybikeshop/apps/warehouse/app/api/middlewares"
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
		"app":     "warehouse",
		"version": warehouse.Version,
	})
}

type ApiDeps struct {
	dig.In

	KratosAPIClient     *ory.APIClient
	InventoryController *inventory.InventoryController
}

func ProvideAPI(deps ApiDeps) *gin.Engine {

	// base API middlewares
	router := ProvideRouter()
	router.Use(gin.Recovery())

	router.Use(middlewares.AuthMiddleware(deps.KratosAPIClient))

	v1 := router.Group("/api/v1")
	deps.InventoryController.Register(v1)

	return router
}
