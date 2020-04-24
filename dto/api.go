package dto

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `form:"username" json:"username" comment:"用户名"  validate:"required" example:""`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:""`
}

func (params *LoginInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ListPageInput struct {
	PageSize int    `form:"page_size" json:"page_size" comment:"每页记录数" validate:"" example:"10"`
	Page     int    `form:"page" json:"page" comment:"页数" validate:"required" example:"1"`
	Name     string `form:"name" json:"name" comment:"姓名" validate:"" example:""`
}

func (params *ListPageInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type AddUserInput struct {
	Name  string `form:"name" json:"name" comment:"姓名" validate:"required"`
	Sex   int    `form:"sex" json:"sex" comment:"性别" validate:""`
	Age   int    `form:"age" json:"age" comment:"年龄" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" comment:"生日" validate:"required"`
	Addr  string `form:"addr" json:"addr" comment:"地址" validate:"required"`
}

func (params *AddUserInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type EditUserInput struct {
	Id    int    `form:"id" json:"id" comment:"ID" validate:"required"`
	Name  string `form:"name" json:"name" comment:"姓名" validate:"required"`
	Sex   int    `form:"sex" json:"sex" comment:"性别" validate:""`
	Age   int    `form:"age" json:"age" comment:"年龄" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" comment:"生日" validate:"required"`
	Addr  string `form:"addr" json:"addr" comment:"地址" validate:"required"`
}

func (params *EditUserInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type RemoveUserInput struct {
	IDS string `form:"ids" json:"ids" comment:"IDS" validate:"required"`
}

func (params *RemoveUserInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
