package router

import (
	"github.com/e421083458/gin_scaffold/controller"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	//demo
	v1 := router.Group("/demo")
	v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware())
	{
		controller.DemoRegister(v1)
	}

	//非登陆接口
	store := sessions.NewCookieStore([]byte("secret"))
	apiNormalGroup := router.Group("/api")
	apiNormalGroup.Use(sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.ApiRegister(apiNormalGroup)
	}

	//登陆接口
	apiAuthGroup := router.Group("/api")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.ApiLoginRegister(apiAuthGroup)
	}
	return router
}
