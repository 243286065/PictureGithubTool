package handler

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	simplejson "github.com/bitly/go-simplejson"
)

func ReadFileContent(filepath string) (string, error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("Read file error: " + err.Error())
		return "", err
	}

	return base64.StdEncoding.EncodeToString(f), nil
}

//UploadFileToGithub: 上传文件到github
func UploadFileToGithub(filename string, filepath string, targetPath string, token string) (string, error) {
	var errMsg string
	client := &http.Client{}
	url := targetPath

	defer os.Remove(filepath)

	fileContent, err := ReadFileContent(filepath)
	if err != nil {
		errMsg = "Cannot read file content"
		log.Println(errMsg + ":" + err.Error())

		return "", err
	}

	body := []byte(fmt.Sprintf("{\"message\":\"update\", \"content\":\"%s\"}", fileContent))
	request, err := http.NewRequest("PUT", url, bytes.NewBufferString(string(body)))

	if err != nil {
		errMsg = "Cannot create new http request"
		log.Println(errMsg + ":" + err.Error())

		return "", err
	}

	request.Header.Add("Authorization", token)

	// log.Println(request)

	response, err := client.Do(request)
	if err != nil {
		errMsg = "Failed to send http request"
		log.Println(errMsg + ": " + err.Error())

		return "", err
	}
	defer response.Body.Close()

	log.Println(fmt.Sprintf("%s : %d", filename, response.StatusCode))

	if response.StatusCode >= 400 {
		return "", errors.New("Invalid status code: " + string(response.StatusCode))
	}

	bodyContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errMsg = "Failed to read http response"
		log.Println(errMsg + ": " + err.Error())
		return "", err
	}

	// log.Println(string(bodyContent))
	js, err := simplejson.NewJson(bodyContent)
	if err != nil {
		errMsg = "Failed to parse http response"
		log.Println(errMsg + ": " + err.Error())
		return "", err
	}
	imagePath, err := js.Get("content").Get("html_url").String()
	if err != nil {
		errMsg = "Failed to get download_url"
		log.Println(errMsg + ": " + err.Error())
		return "", err
	}

	log.Println(imagePath)

	return imagePath + "?raw=true", nil
}
