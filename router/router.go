package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"taibai/common/conf"
	"taibai/controller/user"
	"taibai/controller/welcome"
	"taibai/middleware"
	"time"
)

func GetHttpServer() *gin.Engine {
	router := gin.New()

	file,_ := os.OpenFile(conf.Config.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	c := gin.LoggerConfig{
		Output:file,
		SkipPaths:[]string{"/welcome"},
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s %s \"%s %s %s %d %s \"%s\" %s\"\n",
				params.TimeStamp.Format(time.RFC3339),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
	}
	router.Use(gin.LoggerWithConfig(c))
	router.Use(gin.Recovery())

	router.Use(middleware.GinBodyLogMiddleware)
	router.Use(middleware.MergeParams)

	gUser := router.Group("/user")
	{
		gUser.GET("/login", user.Login)
		gUser.GET("/welogin", user.WeLogin)
		gUser.GET("/logout", welcome.Index)
	}

	gProfile := router.Group("profile")
	{
		gProfile.GET("/info", welcome.Index)
	}

	//router.Use(middleware.Test)
	//r.Use(middleware.GinBodyLogMiddleware)
	//r.Use(middleware.MergeParams)

	return router
}
