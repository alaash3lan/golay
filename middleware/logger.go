package middleware

import (
	_"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


 func Logger () gin.HandlerFunc {
	return func(c *gin.Context) {
		// startTime := time.Now()

		c.Next()
		// duration := time.Since(startTime)

		logrus.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"remoteAddr": c.Request.RemoteAddr,
		}).Info()
	}
 }