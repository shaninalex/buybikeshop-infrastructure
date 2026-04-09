package api

import (
	// "buybikeshop/apps/admin/app/api/middlewares"
	"buybikeshop/apps/admin"
	"buybikeshop/apps/admin/app/api/employee"
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
		"app":     "warehouse",
		"version": admin.Version,
	})
}

type ApiDeps struct {
	dig.In

	KratosAPIClient *ory.APIClient
	Config          *config.Config

	EmployeeController *employee.EmployeeController
}

func ProvideAPI(deps ApiDeps) *gin.Engine {

	// base API middlewares
	router := ProvideRouter()
	router.Use(gin.Recovery())

	// router.Use(middlewares.AuthMiddleware(deps.KratosAPIClient, deps.Config))

	v1 := router.Group("/api/v1/warehouse")
	deps.EmployeeController.Register(v1)

	return router
}
