package middleware

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
)

//设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := public.Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(public.Validate, trans)
			break
		case "en":
			en_translations.RegisterDefaultTranslations(public.Validate, trans)
			break
		case "zh_tw":
			zh_tw_translations.RegisterDefaultTranslations(public.Validate, trans)
			break
		default:
			zh_translations.RegisterDefaultTranslations(public.Validate, trans)
			break
		}

		//自定义错误内容
		//public.Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		//	return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
		//}, func(ut ut.Translator, fe validator.FieldError) string {
		//	t, _ := ut.T("required", fe.Field())
		//	return t
		//})
		
		//设置trans到context
		c.Set("trans", trans)
		c.Next()
	}
}
