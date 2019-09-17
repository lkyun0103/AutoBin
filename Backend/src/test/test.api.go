package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	url := "http://api.choviwu.top/garbage/uploadFile"
	buf := new(bytes.Buffer)	// 图片数据
	writer := multipart.NewWriter(buf)
	writer.WriteField("sublib", "1")
	formFile, err := writer.CreateFormFile("file", "./1.jpg")
	if err != nil {
		fmt.Println("Create form file failed: %s\n", err)
	}
	srcFile, err := os.Open("G:\\git\\AutoBin\\Backend\\src\\waste\\apple.jpg")
	if err != nil {
		fmt.Println("%Open source file failed: s\n", err)
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		fmt.Println("Write to form file falied: %s\n", err)
	}
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	var imageData io.Reader
	imageData.Read(buf.Bytes())
	req, _ := http.NewRequest("POST", url, imageData)
	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	//req.Header.Add("User-Agent", "PostmanRuntime/7.16.3")
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("Cache-Control", "no-cache")
	//req.Header.Add("Postman-Token", "eb1a49ee-99e7-487e-9805-4da100d65ca5,cd993124-9c2f-4b95-bee2-76f030933244")
	//req.Header.Add("Host", "api.choviwu.top")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------153120684342055966849553")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	//req.Header.Add("Content-Length", "4447")
	//req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
