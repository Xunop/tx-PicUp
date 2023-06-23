package main

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	config := Config()
	txUrl := config.Url
	secretID := config.SecretID
	secretKey := config.SecretKey

	if isInputFromPipe() {
		println("Input from pipe")
		u, _ := url.Parse(txUrl)
		b := &cos.BaseURL{BucketURL: u}
		c := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  secretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
				SecretKey: secretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			},
		})
		// 对象键（Key）是对象在存储桶中的唯一标识。
		// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
		name := "test/objectPut.go"
		// 1.通过字符串上传对象
		f := strings.NewReader("test")

		_, err := c.Object.Put(context.Background(), name, f, nil)
		if err != nil {
			panic(err)
		}
		// 2.通过本地文件上传对象
		_, err = c.Object.PutFromFile(context.Background(), name, "../test", nil)
		if err != nil {
			panic(err)
		}
		// 3.通过文件流上传对象
		fd, err := os.Open("./test")
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		_, err = c.Object.Put(context.Background(), name, fd, nil)
		if err != nil {
			panic(err)
		}
	}
}

// Check if input is from pipe
func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
