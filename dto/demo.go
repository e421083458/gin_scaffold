package dto

import (
	"errors"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type InStruct struct {
	Name   string `form:"name" validate:"required"`
	Age    int64  `form:"age" validate:"required"`
	Passwd string `form:"passwd" validate:"required"`
}

func (o *InStruct) BindingValidParams(c *gin.Context)  error{
	if err:=c.ShouldBind(o);err!=nil{
		return err
	}
	v:=c.Value("trans")
	trans,ok := v.(ut.Translator)
	if !ok{
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs:=[]string{}
		for _, e := range errs {
			sliceErrs=append(sliceErrs,e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs,","))
	}
	return nil
}