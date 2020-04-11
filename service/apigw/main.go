// 创建http服务器
package main

import (
	cfg "PictureGithubTool/config"
	"PictureGithubTool/service/apigw/route"
)

func main() {
	router := route.Router()
	router.Run(cfg.WebServerHost)
}
