package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

func AuthMiddleware(kratos *ory.APIClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		session, _, err := kratos.FrontendAPI.ToSession(c.Request.Context()).Cookie(c.GetHeader("Cookie")).Execute()
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		if time.Now().After(*session.ExpiresAt) {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
