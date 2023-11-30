package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/pkg/setting"
	"github.com/xishengcai/cloud/routers"
)

// @title Swagger Example API
// @version 2.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	klog.InitFlags(flag.CommandLine)
	flag.Parse()
	gin.SetMode(setting.Config.RunMode)
	//db.InitMysql()

	handler := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", setting.Config.Web.Port),
		Handler:        handler,
		ReadTimeout:    setting.Config.Web.ReadTimeout * time.Second,
		WriteTimeout:   setting.Config.Web.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	//Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	_ = server.Close()

}
