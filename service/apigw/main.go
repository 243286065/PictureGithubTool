package main

import (
	cfg "PictureGithubTool/config"
	"PictureGithubTool/service/apigw/route"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	router := route.Router()
	router.Run(cfg.WebServerHost)
}
