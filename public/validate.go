package public

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
)
import "gopkg.in/go-playground/validator.v9"

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InitValidate()  {
	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()
}

