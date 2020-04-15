package dto

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

type DemoInput struct {
	Name   string `form:"name" comment:"姓名" validate:"required"`
	Age    int64  `form:"age" comment:"年龄" validate:"required"`
	Passwd string `form:"passwd" comment:"密码" validate:"required"`
}

func (params *DemoInput) BindingValidParams(c *gin.Context)  error{
	return public.DefaultGetValidParams(c, params)
}