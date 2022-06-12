package main

import (
	"douyin/controller"
	"douyin/repository"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()
	initRouter(r)
	r.Run(controller.Port)
}

func Init() error {

	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	if err := util.CreateDirIfNotExist("./public"); err != nil {
		return err
	}
	return nil
}
