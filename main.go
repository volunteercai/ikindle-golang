package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "ikindle-golang/src/api"
	"ikindle-golang/src/common"
	"strconv"
)

func init() {
	fmt.Println("初始化")
}

func main() {
	loadConfig()

	engine := gin.Default()
	loadRoutes(engine)

	err := engine.Run(":" + strconv.Itoa(common.GetConfig().Web.Port))
	panic(err)
}

func loadRoutes(engine *gin.Engine) {
	for _, route := range common.Routes {
		group := engine.Group(route.Path)
		for _, childRoute := range route.Children {
			group.Handle(childRoute.HttpMethod, childRoute.Path, childRoute.Handler)
		}
	}
}

func loadConfig() {
	common.LoadConfig("resources/config/config.yml")
}
