package controller

import (
	"encoding/json"
	"github.com/e421083458/gin_scaffold/dao"
	"github.com/e421083458/gin_scaffold/dto"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
}

func DemoRegister(router *gin.RouterGroup) {
	demo := DemoController{}
	router.GET("/index", demo.Index)
	router.GET("/bind", demo.Bind)
	router.GET("/dao", demo.Dao)
	router.GET("/redis", demo.Redis)
}

func (demo *DemoController) Index(c *gin.Context) {
	middleware.ResponseSuccess(c, "")
	return
}

func (demo *DemoController) Dao(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	area, err := (&dao.Area{}).Find(c, tx, c.DefaultQuery("id", "1"));
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	js, _ := json.Marshal(area)
	middleware.ResponseSuccess(c, string(js))
	return
}

func (demo *DemoController) Redis(c *gin.Context) {
	redisKey := "redis_key"
	lib.RedisConfDo(public.GetTraceContext(c),
		"default",
		"SET",
		redisKey, "redis_value")
	redisValue, err := redis.String(
		lib.RedisConfDo(public.GetTraceContext(c), "default",
		"GET",
		redisKey))
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	middleware.ResponseSuccess(c, redisValue)
	return
}

func (demo *DemoController) Bind(c *gin.Context) {
	st := &dto.DemoInput{}
	if err := st.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
