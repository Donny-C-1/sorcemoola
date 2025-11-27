package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	origin := c.GetHeader("Origin")

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ";")
	fmt.Println(allowedOrigins)

	allowOrigin := ""
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			allowOrigin = origin
			break
		}
	}

	if allowOrigin != "" {
		c.Header("Access-Control-Allow-Origin", allowOrigin)
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, X-CSRF-Token, Cache-Control, X-Requested-With")

	c.Next()
}
