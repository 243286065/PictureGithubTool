package route

import (
	"PictureGithubTool/service/apigw/handler"

	"github.com/gin-gonic/gin"
)

// Router 返回路由器
func Router() *gin.Engine {
	//初始化
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()

	//处理静态资源
	router.Static("/static/", "./static")

	//处理默认首页
	router.GET("/", handler.DefaultHomePageHandler)

	return router
}
