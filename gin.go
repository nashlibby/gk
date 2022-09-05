package gk

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

// gin 成功响应
func GinSuccess(c *gin.Context, code int, message string, data ...interface{}) {
	var d any
	if len(data) == 0 {
		d = nil
	} else {
		d = data[0]
	}

	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    d,
	})
	c.Abort()
	return
}

// gin 失败响应
func GinFail(c *gin.Context, code int, message string) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    nil,
	})
	c.Abort()
	return
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// gin 日志中间件
func GinLogger() gin.HandlerFunc {
	logger := Logger("api")
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		// 开始时间
		startTime := time.Now()
		// 请求body
		requestBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 响应body
		responseBody := blw.body.String()
		//日志格式
		logger.Infof("| %3d | %13v | %12s | %s | %s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			string(requestBody),
			responseBody,
		)
	}
}
