package routers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/logto-io/go/client"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/xishengcai/cloud/docs"
	"github.com/xishengcai/cloud/pkg/middleware"
	"github.com/xishengcai/cloud/routers/api/cluster"
	"github.com/xishengcai/cloud/routers/api/image"
	"github.com/xishengcai/cloud/routers/api/proxy"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	// We use memory-based session in this example
	store := memstore.NewStore([]byte("your session secret"))
	r.Use(sessions.Sessions("session", store))
	r.Use(gin.Logger())
	r.Use(Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	logtoConfig := &client.LogtoConfig{
		Endpoint:  "http://localhost:3001/",
		AppId:     "yei5skxna17gu4ljzilhu",
		AppSecret: "KjA61mL33yEQ90Xt4aEMPxeYL8r0KlJB",
	}
	// Use Logto to control the content of the home page
	authState := "You are not logged in to this website. :("

	// Add a link to perform a sign-in request on the home page
	r.GET("/", func(ctx *gin.Context) {
		// ...
		homePage := `<h1>Hello Logto</h1>` +
			"<div>" + authState + "</div>" +
			`<div><a href="/sign-in">Sign In</a></div>` +
			// Add link
			`<div><a href="/sign-out">Sign Out</a></div>`

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(homePage))
	})

	// Add a route for handling sign-in requests
	r.GET("/sign-in", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(
			logtoConfig,
			&SessionStorage{session: session},
		)

		// The sign-in request is handled by Logto.
		// The user will be redirected to the Redirect URI on signed in.
		signInUri, err := logtoClient.SignIn("http://localhost/sign-in-callback")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Redirect the user to the Logto sign-in page.
		ctx.Redirect(http.StatusTemporaryRedirect, signInUri)
	})

	// Add a route for handling sign-in callback requests
	r.GET("/sign-in-callback", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(
			logtoConfig,
			&SessionStorage{session: session},
		)

		// The sign-in callback request is handled by Logto
		err := logtoClient.HandleSignInCallback(ctx.Request)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Jump to the page specified by the developer.
		// This example takes the user back to the home page.
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	// Add a route for handling signing out requests
	r.GET("/sign-out", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(
			logtoConfig,
			&SessionStorage{session: session},
		)

		// The sign-out request is handled by Logto.
		// The user will be redirected to the Post Sign-out Redirect URI on signed out.
		signOutUri, signOutErr := logtoClient.SignOut("http://localhost")

		if signOutErr != nil {
			ctx.String(http.StatusOK, signOutErr.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, signOutUri)
	})

	baseRoute := r.Group("/api/v1")
	//baseRoute.Use(Auth())
	cluster.Register(baseRoute)
	proxy.Register(baseRoute)
	image.Register(baseRoute)
	agg := r.Group("/")
	agg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
