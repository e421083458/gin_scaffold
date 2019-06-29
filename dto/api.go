package dto

import (
	"errors"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type LoginInput struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (o *LoginInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

type ListPageInput struct {
	Page string `form:"page" json:"page" validate:"required"`
	Name string `form:"name" json:"name" validate:""`
}

func (o *ListPageInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

type AddUserInput struct {
	Name  string `form:"name" json:"name" validate:"required"`
	Sex   int    `form:"sex" json:"sex" validate:""`
	Age   int    `form:"age" json:"age" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" validate:"required"`
	Addr  string `form:"addr" json:"addr" validate:"required"`
}

func (o *AddUserInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

type EditUserInput struct {
	Id    int `form:"id" json:"id" validate:"required"`
	Name  string `form:"name" json:"name" validate:"required"`
	Sex   int    `form:"sex" json:"sex" validate:""`
	Age   int    `form:"age" json:"age" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" validate:"required"`
	Addr  string `form:"addr" json:"addr" validate:"required"`
}

func (o *EditUserInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}


type RemoveUserInput struct {
	IDS    string `form:"ids" json:"ids" validate:"required"`
}

func (o *RemoveUserInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
