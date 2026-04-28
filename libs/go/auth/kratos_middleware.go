package auth

import (
	"buybikeshop/libs/go/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

func AuthMiddleware(kratos *ory.APIClient, cnf *config.Config) func(c *gin.Context) {
	return func(c *gin.Context) {
		session, _, err := kratos.FrontendAPI.ToSession(c.Request.Context()).Cookie(c.GetHeader("Cookie")).Execute()
		if err != nil {
			c.Status(http.StatusUnauthorized)
			log.Printf("[AuthMiddleware]: %s\n", err)
			c.Abort()
			return
		}

		if time.Now().After(*session.ExpiresAt) {
			log.Println("[AuthMiddleware]: expired")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set(config.SessionId, session.GetId())
		c.Set(config.IdentityId, session.Identity.GetId())

		c.Next()
	}
}
