package main

import (
	"github.com/e421083458/gin_scaffold/router"
	"github.com/e421083458/golang_common/lib"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",})
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}