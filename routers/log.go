package routers

import (
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

var (
	skipPaths = []string{
		"/health",
		"/swagger/.*",
	}

	skipHandlerFunMap = make(map[string]struct{})
)

func matchPath(path string) bool {
	for _, i := range skipPaths {
		if i == path {
			return true
		}
		if regexp.MustCompile(i).Match([]byte(path)) {
			return true
		}
	}
	return false
}

func getSkipHandlerFunMap() map[string]struct{} {
	return skipHandlerFunMap
}

// Logger 用于gin请求调用时，输出请求体，响应体的日志中间件
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			requestBody          []byte
			bodyStr              string
			buf                  = new(bytes.Buffer)
			customResponseWriter = &customResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		)
		if matchPath(ctx.Request.URL.Path) {
			return
		}
		if ctx.Request.Body != nil {
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			bodyStr = string(requestBody)
		}

		//ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
		ctx.Request.Body = &seekerReaderCloser{bytes.NewReader(requestBody)}
		ctx.Writer = customResponseWriter

		//强制打印颜色（默认Output为终端输出时才会打印）
		gin.ForceConsoleColor()
		gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: nil,
			Output:    buf,
		})(ctx)

		log := buf.String()
		if log == "" {
			return
		}

		handlerFunMap := getSkipHandlerFunMap()
		if len(handlerFunMap) > 0 {
			name := ctx.HandlerName()
			if _, ok := handlerFunMap[name]; ok {
				fmt.Print(log)
				return
			}
		}
		klog.V(2).Infof("Request Header: %v", ctx.Request.Header)
		klog.V(1).Infof("Request Body: %s", bodyStr)
		responseStr := customResponseWriter.body.String()
		klog.V(1).Infof("Response Body: %s", responseStr)
	}
}

type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w customResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w customResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

type seekerReaderCloser struct {
	*bytes.Reader
}

func (seekerReaderCloser) Close() error { return nil }
