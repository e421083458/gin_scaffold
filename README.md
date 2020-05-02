<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [gin_scaffold](#gin_scaffold)
    - [现在开始](#%E7%8E%B0%E5%9C%A8%E5%BC%80%E5%A7%8B)
    - [文件分层](#%E6%96%87%E4%BB%B6%E5%88%86%E5%B1%82)
    - [log / redis / mysql / http.client 常用方法](#log--redis--mysql--httpclient-%E5%B8%B8%E7%94%A8%E6%96%B9%E6%B3%95)
    - [swagger文档生成](#swagger%E6%96%87%E6%A1%A3%E7%94%9F%E6%88%90)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# gin_scaffold
Gin best practices, gin development scaffolding, too late to explain, get on the bus.

使用gin构建了企业级脚手架，代码简洁易读，可快速进行高效web开发。
主要功能有：
1. 请求链路日志打印，涵盖mysql/redis/request
2. 支持多语言错误信息提示及自定义错误提示。
3. 支持了多配置环境
4. 封装了 log/redis/mysql/http.client 常用方法
5. 支持swagger文档生成

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
- 确保正确配置了 conf/mysql_map.toml、conf/redis_map.toml：

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

创建测试表：
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
    "errmsg": "年龄为必填字段,密码为必填字段",
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
├── conf            配置文件夹
│   └── dev
│       ├── base.toml
│       ├── mysql_map.toml
│       └── redis_map.toml
├── controller      控制器
│   └── demo.go
├── dao             DB数据层
│   └── demo.go
├── docs            swagger文件层
├── dto             输入输出结构层
│   └── demo.go
├── go.mod
├── go.sum
├── main.go         入口文件
├── middleware      中间件层
│   ├── panic.go
│   ├── response.go
│   ├── token_auth.go
│   └── translation.go
├── public          公共文件
│   ├── log.go
│   ├── mysql.go
│   └── validate.go
└── router          路由层
    ├── httpserver.go
    └── route.go
```

### log / redis / mysql / http.client 常用方法

参考文档：https://github.com/e421083458/golang_common


### swagger文档生成

https://github.com/swaggo/swag/releases

- 下载对应操作系统的执行文件到$GOPATH/bin下面

如下：
```
➜  gin_scaffold git:(master) ✗ ll -r $GOPATH/bin
total 434168
-rwxr-xr-x  1 niuyufu  staff    13M  4  3 17:38 swag
```

- 设置接口文档参考： `controller/demo.go` 的 Bind方法的注释设置

```
// ListPage godoc
// @Summary 测试数据绑定
// @Description 测试数据绑定
// @Tags 用户
// @ID /demo/bind
// @Accept  json
// @Produce  json
// @Param polygon body dto.DemoInput true "body"
// @Success 200 {object} middleware.Response{data=dto.DemoInput} "success"
// @Router /demo/bind [post]
```

- 生成接口文档：`swag init`
- 然后启动服务器：`go run main.go`，浏览地址: http://127.0.0.1:8880/swagger/index.html