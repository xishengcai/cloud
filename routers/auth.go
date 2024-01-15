package routers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/logto-io/go/client"

	"github.com/xishengcai/cloud/pkg/app"
)

type SessionStorage struct {
	session sessions.Session
}

func (storage *SessionStorage) GetItem(key string) string {
	value := storage.session.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func (storage *SessionStorage) SetItem(key, value string) {
	storage.session.Set(key, value)
	storage.session.Save()
}

func auth(ctx *gin.Context) {
	logtoConfig := &client.LogtoConfig{
		Endpoint:  "http://localhost:3001/",
		AppId:     "yei5skxna17gu4ljzilhu",
		AppSecret: "KjA61mL33yEQ90Xt4aEMPxeYL8r0KlJB",
	}
	session := sessions.Default(ctx)
	sessionStorage := &SessionStorage{session: session}

	logtoClient := client.NewLogtoClient(
		logtoConfig,
		sessionStorage,
	)
	if !logtoClient.IsAuthenticated() {
		ctx.JSON(http.StatusUnauthorized, app.Response{
			Code:   http.StatusUnauthorized,
			ResMsg: "unauthorized",
		})
		ctx.Abort()
	}
}

func Auth() gin.HandlerFunc {
	return auth
}
