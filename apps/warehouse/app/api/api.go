package api

import (
	"buybikeshop/apps/warehouse"
	"buybikeshop/apps/warehouse/app/api/controllers/inventory"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func ProvideRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

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

	InventoryController *inventory.InventoryController
}

func NewApi(deps ApiDeps) *gin.Engine {

	// base API middlewares
	router := ProvideRouter()
	router.Use(gin.Recovery()) // <= write your own recovery middleware

	// router.Use(middlewares.LoggingMiddleware())
	// router.Use(middlewares.CORSMiddleware())

	v1 := router.Group("/api/v1")
	deps.InventoryController.Register(v1)

	return router
}
