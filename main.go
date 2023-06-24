package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

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
				SecretID:  secretID,  // https://cloud.tencent.com/document/product/598/37140
				SecretKey: secretKey, // https://cloud.tencent.com/document/product/598/37140
			},
		})
		// 对象键（Key）是对象在存储桶中的唯一标识。
		// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
		// 这里设置为当前时间
		now := time.Now().Unix()
		name := fmt.Sprintf("notes/PicUp-%d.png", now)
		// 3.通过文件流上传对象
		if isInputFromPipe() {
			_, err := c.Object.Put(context.Background(), name, os.Stdin, nil)
			if err != nil {
				panic(err)
			}
			// 输出上传的文件地址 方便复制
			os.Stdout.WriteString(txUrl + "/" + name)
		}
	}
}

// Check if input is from pipe
func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
