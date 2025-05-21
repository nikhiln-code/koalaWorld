package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger returns a Gin middleware that logs requests using Zap.
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next() // process request

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		clientIP := c.ClientIP()
		errorMsg := c.Errors.ByType(gin.ErrorTypePrivate).String()

		Log().Info("HTTP Request",
			zap.String("method", method),
			zap.String("path", path),
			zap.String("query", raw),
			zap.Int("status", status),
			zap.String("ip", clientIP),
			zap.String("latency", latency.String()),
			zap.String("error", errorMsg),
		)
	}
}

func GinRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(zap.NewStdLog(Log()).Writer())
}
