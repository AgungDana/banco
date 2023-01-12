package ctxutil

import (
	"banco/common/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func GinAuth(c *gin.Context) {
	ctx := c.Request.Context()

	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatus(403)
		return
	}
	newHeader := strings.Split(header, " ")
	if newHeader[0] != "Bearer" {
		c.JSON(403, gin.H{
			"error": "invalid token",
		})
		c.Abort()
		return
	}
	if newHeader[1] == "" {
		c.JSON(403, gin.H{
			"error": "need authorization",
		})
		c.Abort()
		return
	}

	token := jwt.NewJwtToken("1234567890")
	payload, err := token.ValidateToken(newHeader[1])

	if err != nil {
		c.JSON(403, gin.H{
			"error": err,
		})
		c.Abort()
		return
	}
	ctx = NewRequestPayload(ctx, *payload)
	ctx = NewRequestUserID(ctx, payload.UserID)

	req := c.Request.WithContext(ctx)
	c.Request = req
}

func SetTransaction(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = NewRequestTransactionId(ctx)
	req := c.Request.WithContext(ctx)
	c.Request = req
}
