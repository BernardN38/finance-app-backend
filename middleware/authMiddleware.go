package middleware

import (
	"fmt"

	"github.com/bernardn38/financefirst/token"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(authorizationPayloadKey)
		if err != nil {
			c.AbortWithStatusJSON(401, "no auth cookie provided")
			return
		}

		payload, err := tokenMaker.VerifyToken(cookie)
		fmt.Println(payload)
		if err != nil {
			c.AbortWithStatusJSON(401, "cookie cant be verified")
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
