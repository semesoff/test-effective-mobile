package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Since(startTime)
		logrus.Debug(fmt.Sprintf("Request: %s Duration: %d.03 ms", c.Request.URL.Path, endTime.Milliseconds()))
	}
}
