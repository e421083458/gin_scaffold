package router

import (
	"github.com/e421083458/gin_scaffold/controller"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	//写入gin日志
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter = io.MultiWriter(f)
	router := gin.Default()
	router.Use(middlewares...)

	//demo
	v1 := router.Group("/demo")
	v1.Use(middleware.RecoveryMiddleware(), middleware.TokenAuthMiddleware(), middleware.TranslationMiddleware() )
	{
		controller.DemoRegister(v1)
	}
	return router
}