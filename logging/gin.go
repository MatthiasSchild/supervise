package logging

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func DevGinLogger(c *gin.Context) {
	accessCounter := GetAccessCounter(c)
	start := time.Now()
	c.Next()

	msg := fmt.Sprintf(
		"%03d | %s %s",
		c.Writer.Status(),
		strings.ToUpper(c.Request.Method),
		c.Request.URL.String(),
	)

	slog.Info(
		msg,
		"access", accessCounter,
		"httpRequest", map[string]any{
			"requestMethod": c.Request.Method,
			"requestUrl":    c.Request.URL.String(),
			"remoteIp":      c.ClientIP(),
			"status":        c.Writer.Status(),
			"latency":       float64(time.Since(start).Microseconds()) / 1000.0,
		},
	)
}

func ProdGinLogger(c *gin.Context) {
	accessCounter := GetAccessCounter(c)
	start := time.Now()
	c.Next()

	slog.Info(
		"http request",
		"access", accessCounter,
		"httpRequest", map[string]any{
			"requestMethod": c.Request.Method,
			"requestUrl":    c.Request.URL.String(),
			"remoteIp":      c.ClientIP(),
			"status":        c.Writer.Status(),
			"latency":       float64(time.Since(start).Microseconds()) / 1000.0,
		},
	)
}

func GinLoggerByEnv(variableName string, expected string) gin.HandlerFunc {
	if os.Getenv(variableName) == expected {
		return ProdGinLogger
	} else {
		return DevGinLogger
	}
}