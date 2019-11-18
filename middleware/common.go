package middleware

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Common(c *gin.Context) {
	r := c.Request
	reqID := c.Request.Header.Get("request_id")
	if reqID == "" {
		reqID = GenLogId()
	}
	startTime := time.Now()
	logBuf := log.WithFields(log.Fields{
		"path":       r.URL.Path,
		"method":     r.Method,
		"request_id": reqID,
		"client_ip":  r.RemoteAddr,
		"start_time": startTime,
	})

	c.Set("logger", logBuf)
	c.Set("request_id", reqID)
	c.Set("start_time", startTime)

	c.Next()
}

//依据时间戳产生logid，该logid方便后续根据真实时间戳分类存储图文日志
func GenLogId() string {
	return fmt.Sprintf("%d.%d", time.Now().Unix(), rand.Intn(10000))
}
