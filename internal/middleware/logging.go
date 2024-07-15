package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		startTime := time.Now()

		ctx.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := ctx.Request.Method

		reqUri := ctx.Request.RequestURI

		statusCode := ctx.Writer.Status()

		clientIP := ctx.ClientIP()

		log.WithFields(map[string]interface{}{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Info("HTTP REQUEST")

		ctx.Next()
	}
}
