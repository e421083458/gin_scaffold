package public

import (
	"context"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
)

//错误日志
func ContextWarning(c context.Context, dltag string, m map[string]interface{}) {
	v:=c.Value("trace")
	traceContext,ok := v.(*lib.TraceContext)
	if !ok{
		traceContext = lib.NewTrace()
	}
	lib.Log.TagWarn(traceContext, dltag, m)
}

//错误日志
func ContextError(c context.Context, dltag string, m map[string]interface{}) {
	v:=c.Value("trace")
	traceContext,ok := v.(*lib.TraceContext)
	if !ok{
		traceContext = lib.NewTrace()
	}
	lib.Log.TagError(traceContext, dltag, m)
}

//普通日志
func ContextNotice(c context.Context, dltag string, m map[string]interface{}) {
	v:=c.Value("trace")
	traceContext,ok := v.(*lib.TraceContext)
	if !ok{
		traceContext = lib.NewTrace()
	}
	lib.Log.TagInfo(traceContext, dltag, m)
}

//错误日志
func ComLogWarning(c *gin.Context, dltag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagError(traceContext, dltag, m)
}

//普通日志
func ComLogNotice(c *gin.Context, dltag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagInfo(traceContext, dltag, m)
}

// 从gin的Context中获取数据
func GetGinTraceContext(c *gin.Context) *lib.TraceContext {
	// 防御
	if c == nil {
		return lib.NewTrace()
	}
	traceContext, exists := c.Get("trace")
	if exists {
		if tc, ok := traceContext.(*lib.TraceContext); ok {
			return tc
		}
	}
	return lib.NewTrace()
}

// 从Context中获取数据
func GetTraceContext(c context.Context) *lib.TraceContext {
	if c == nil {
		return lib.NewTrace()
	}
	traceContext:=c.Value("trace")
	if tc, ok := traceContext.(*lib.TraceContext); ok {
		return tc
	}
	return lib.NewTrace()
}
