# gin_scaffold
Gin best practices, gin development scaffolding, too late to explain, get on the bus.

使用gin构建了企业级脚手架，代码简洁易读，可快速进行高效web开发。
主要功能有：
1. 请求链路日志打印，涵盖mysql/redis/request_in/request_out
2. 接入validator.v9，支持多语言错误信息提示及自定义错误提示。
3. 借助golang_common，支持了多配置环境及log/redis/mysql/http.client

项目地址：https://github.com/e421083458/gin_scaffold
### 现在开始
- 安装软件依赖
go mod使用请查阅：

https://blog.csdn.net/e421083458/article/details/89762113
```
git clone git@github.com:e421083458/gin_scaffold.git
cd gin_scaffold
go mod tidy
```
- 运行脚本
```
go run main.go

➜  gin_scaffold git:(master) ✗ go run main.go
------------------------------------------------------------------------
[INFO]  config=./conf/dev/
[INFO]  start loading resources.
[INFO]  success loading resources.
------------------------------------------------------------------------
[GIN-debug] [WARNING] Now Gin requires Go 1.6 or later and Go 1.7 will be required soon.

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /demo/index               --> github.com/e421083458/gin_scaffold/controller.(*Demo).Index-fm (6 handlers)
[GIN-debug] GET    /demo/bind                --> github.com/e421083458/gin_scaffold/controller.(*Demo).Bind-fm (6 handlers)
[GIN-debug] GET    /demo/dao                 --> github.com/e421083458/gin_scaffold/controller.(*Demo).Dao-fm (6 handlers)
[GIN-debug] GET    /demo/redis               --> github.com/e421083458/gin_scaffold/controller.(*Demo).Redis-fm (6 handlers)
 [INFO] HttpServerRun::8880
```
- 测试mysql与请求链路

创建测试表,并确保正确配置 mysql_map.toml：
```
CREATE TABLE `area` (
 `id` bigint(20) NOT NULL AUTO_INCREMENT,
 `area_name` varchar(255) NOT NULL,
 `city_id` int(11) NOT NULL,
 `user_id` int(11) NOT NULL,
 `update_at` datetime NOT NULL,
 `create_at` datetime NOT NULL,
 `delete_at` datetime NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='area';
INSERT INTO `area` (`id`, `area_name`, `city_id`, `user_id`, `update_at`, `create_at`, `delete_at`) VALUES (NULL, 'area_name', '1', '2', '2019-06-15 00:00:00', '2019-06-15 00:00:00', '2019-06-15 00:00:00');
```


```
curl 'http://127.0.0.1:8880/demo/dao?id=1'
{
    "errno": 0,
    "errmsg": "",
    "data": "[{\"id\":1,\"area_name\":\"area_name\",\"city_id\":1,\"user_id\":2,\"update_at\":\"2019-06-15T00:00:00+08:00\",\"create_at\":\"2019-06-15T00:00:00+08:00\",\"delete_at\":\"2019-06-15T00:00:00+08:00\"}]",
    "trace_id": "c0a8fe445d05b9eeee780f9f5a8581b0"
}

查看链路日志（确认是不是一次请求查询，都带有相同trace_id）：
tail -f gin_scaffold.inf.log

[INFO][2019-06-16T11:39:26.802][log.go:58] _com_request_in||method=GET||from=127.0.0.1||traceid=c0a8fe445d05b9eeee780f9f5a8581b0||cspanid=||uri=/demo/dao?id=1||args=map[]||body=||spanid=9dad47aa57e9d186
[INFO][2019-06-16T11:39:26.802][log.go:58] _com_mysql_success||affected_row=1||traceid=c0a8fe445d05b9ee07b80f9f66cb39b0||spanid=9dad47aa1408d2ac||source=/Users/niuyufu/go/src/github.com/e421083458/gin_scaffold/dao/demo.go:24||proc_time=0.000000000||sql=SELECT * FROM `area`  WHERE (id = '1')||level=sql||current_time=2019-06-16 11:39:26||cspanid=
[INFO][2019-06-16T11:39:26.802][log.go:58] _com_request_out||method=GET||args=map[]||proc_time=0.025019164||traceid=c0a8fe445d05b9eeee780f9f5a8581b0||spanid=9dad47aa57e9d186||uri=/demo/dao?id=1||from=127.0.0.1||response={\"errno\":0,\"errmsg\":\"\",\"data\":\"[{\\\"id\\\":1,\\\"area_name\\\":\\\"area_name\\\",\\\"city_id\\\":1,\\\"user_id\\\":2,\\\"update_at\\\":\\\"2019-06-15T00:00:00+08:00\\\",\\\"create_at\\\":\\\"2019-06-15T00:00:00+08:00\\\",\\\"delete_at\\\":\\\"2019-06-15T00:00:00+08:00\\\"}]\",\"trace_id\":\"c0a8fe445d05b9eeee780f9f5a8581b0\"}||cspanid=
```
- 测试参数绑定与多语言验证

```
curl 'http://127.0.0.1:8880/demo/bind?name=name&locale=zh'
{
    "errno": 500,
    "errmsg": "Age为必填字段,Passwd为必填字段",
    "data": "",
    "trace_id": "c0a8fe445d05badae8c00f9fb62158b0"
}

curl 'http://127.0.0.1:8880/demo/bind?name=name&locale=en'
{
    "errno": 500,
    "errmsg": "Age is a required field,Passwd is a required field",
    "data": "",
    "trace_id": "c0a8fe445d05bb4cd3b00f9f3a768bb0"
}
```

### 文件分层
```
├── README.md
├── conf   配置文件夹
│   └── dev
│       ├── base.toml
│       ├── mysql_map.toml
│       └── redis_map.toml
├── controller 控制器
│   └── demo.go
├── dao DB数据访问层
│   └── demo.go
├── dto  Bind结构体层
│   └── demo.go
├── gin_scaffold.inf.log  info日志
├── gin_scaffold.wf.log warning日志
├── go.mod go module管理文件
├── go.sum
├── main.go
├── middleware 中间件层
│   ├── panic.go
│   ├── response.go
│   ├── token_auth.go
│   └── translation.go
├── public  公共文件
│   ├── log.go
│   ├── mysql.go
│   └── validate.go
├── router  路由层
│   ├── httpserver.go
│   └── route.go
└── tmpl
```
### 引入轻量级golang类库，支持 mysql、redis、http.client、log、支持多级环境配置、支持链路日志打印
```
go get -v github.com/e421083458/golang_common
```
> 测试日志打印功能：
```
package main

import (
	"github.com/e421083458/golang_common/lib"
	"log"
	"time"
)

func main() {
	if err := lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",}); err != nil {
		log.Fatal(err)
	}
	defer lib.Destroy()

	//todo sth
	lib.Log.TagInfo(lib.NewTrace(), lib.DLTagUndefind, map[string]interface{}{"message": "todo sth"})
	time.Sleep(time.Second)
}
```
类库更多功能细节请查看：
https://github.com/e421083458/golang_common

### 输出格式统一封装
```
func ResponseError(c *gin.Context, code ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "", TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
```
### 定义中间件链路日志打印
```
package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)
//链路请求日志
func RequestInLog(c *gin.Context) {
	traceContext := lib.NewTrace()
	if traceId := c.Request.Header.Get("com-header-rid"); traceId != "" {
		traceContext.TraceId = traceId
	}
	if spanId := c.Request.Header.Get("com-header-spanid"); spanId != "" {
		traceContext.SpanId = spanId
	}
	c.Set("startExecTime", time.Now())
	c.Set("trace", traceContext)
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back

	lib.Log.TagInfo(traceContext, "_com_request_in", map[string]interface{}{
		"uri":    c.Request.RequestURI,
		"method": c.Request.Method,
		"args":   c.Request.PostForm,
		"body":   string(bodyBytes),
		"from":   c.ClientIP(),
	})
}
//链路输出日志
func RequestOutLog(c *gin.Context) {
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")
	startExecTime, _ := st.(time.Time)
	public.ComLogNotice(c, "_com_request_out", map[string]interface{}{
		"uri":       c.Request.RequestURI,
		"method":    c.Request.Method,
		"args":      c.Request.PostForm,
		"from":      c.ClientIP(),
		"response":  response,
		"proc_time": endExecTime.Sub(startExecTime).Seconds(),
	})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestInLog(c)
		defer RequestOutLog(c)
		isMatched := false
		for _, host := range lib.GetStringSliceConf("base.http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched{
			ResponseError(c, InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
```
### 请求数据绑定到结构体与校验

dto/demo.go
```
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
```
controller/demo.go
```
func (demo *Demo) Bind(c *gin.Context) {
	st:=&dto.InStruct{}
	if err:=st.BindingValidParams(c);err!=nil{
		middleware.ResponseError(c,500,err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
```