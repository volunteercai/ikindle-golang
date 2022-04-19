package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ikindle-golang/src/common"
	"ikindle-golang/src/dao"
	"net/http"
)

type UserApi struct {
	common.Api
}

func init() {
	fmt.Println("初始化UserApi")
	route := &common.Route{Path: "/users"}
	api := UserApi{Api: common.Api{R: route}}

	api.LoadRoute()
}

func (r *UserApi) LoadRoute() {
	common.AddRootRoute(r.R)
	r.getUserList()
	r.saveUser()
	r.deleteUser()
}

func (r *UserApi) getUserList() {
	r.R.AddChildRoute(http.MethodGet, "/list", func(context *gin.Context) {
		context.JSON(200, []dao.User{{Name: "蔡大哥哥"}})
	})
}

func (r *UserApi) saveUser() {
	r.R.AddChildRoute(http.MethodPost, "/save", func(context *gin.Context) {
		context.JSON(200, []dao.User{{Name: "蔡大哥哥"}})
	})
}

func (r *UserApi) deleteUser() {
	r.R.AddChildRoute(http.MethodPost, "/delete", func(context *gin.Context) {
		context.JSON(200, []dao.User{{Name: "蔡大哥哥"}})
	})
}
