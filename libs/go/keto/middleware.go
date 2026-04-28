package keto

import (
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/transport"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(p PermissionCheck, object, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if p.Check(c.Request.Context(), config.MustIdentityId(c), object, action) {
			transport.Error(c, http.StatusForbidden, errors.New("not allowed"))
			c.Abort()
			return
		}

		c.Next()
	}
}
