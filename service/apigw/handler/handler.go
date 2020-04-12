package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// "github.com/micro/go-micro"

//初始化微服务，本项目用不上
// func init() {
// 	service := micro.NewService()
// 	// 初始化，解析命令行参数等
// 	service.Init()
// }

//DefaultHomePageHandler 处理默认首页
func DefaultHomePageHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/index.html")
}
