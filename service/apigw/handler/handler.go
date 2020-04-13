package handler

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	cfg "PictureGithubTool/config"

	"github.com/gin-gonic/gin"
)

// "github.com/micro/go-micro"

//初始化微服务，本项目用不上
func init() {
	// service := micro.NewService()
	// // 初始化，解析命令行参数等
	// service.Init()

	//初始化临时目录
	err := os.MkdirAll(cfg.TmpStoreDir, 0777)
	if err != nil {
		log.Println(err.Error())
	}
}

//DefaultHomePageHandler: 处理默认首页
func DefaultHomePageHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/index.html")
}

//UploadFileHandler: 处理文件上传请求
func UploadFileHandler(c *gin.Context) {
	var errMsg string
	err := c.Request.ParseMultipartForm(128)

	if err != nil {
		errMsg = "Cannot parse file upload request"
		log.Println("Cannot parse file upload request: " + err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": errMsg,
		})

		return
	}

	filesFile := c.Request.MultipartForm.File
	files := make(map[string][]*multipart.FileHeader)

	for key, value := range filesFile {
		files[key] = value
	}

	//先将客户端上传的文件上传到临时目录
	for _, headers := range files {
		var errMsg string
		for _, header := range headers {
			file, err := header.Open()
			if err != nil {
				errMsg = "annot parse file header"
				log.Println(errMsg + ": " + err.Error())

				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": errMsg,
				})
				return
			}

			path := cfg.TmpStoreDir + "/" + header.Filename

			dst, err := os.Create(path)
			defer dst.Close()

			if err != nil {
				errMsg = "Cannot create file"
				log.Println(errMsg + ": " + err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": errMsg,
				})

				return
			}

			_, err = io.Copy(dst, file)

			if err != nil {
				errMsg = "Cannot copy file"
				log.Println(errMsg + ": " + err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": errMsg,
				})

				return
			}

			//获取头部信息
			targetPath := c.Request.Header.Get("Target-Path")
			token := c.Request.Header.Get("Authorization")

			//保存文件成功后，上传到github
			imagePath, err := UploadFileToGithub(header.Filename, path, targetPath, token)
			if err != nil {
				errMsg = "Failed to upload file to github"
				log.Println(errMsg + ": " + err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": errMsg,
				})

				return
			}

			//上传成功,返回文件在服务器上的地址
			c.JSON(http.StatusOK, gin.H{
				"imagePath": imagePath,
			})
		}
	}
}
