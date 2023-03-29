package service

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"stouch_server/src/core"
	"strings"
)

var bucket *oss.Bucket

func init_copy() {
	// 创建OSSClient实例。
	client, err := oss.New(core.Config.Oss.EndPoint, core.Config.Oss.AccessKeyId, core.Config.Oss.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err = client.Bucket(core.Config.Oss.BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 上传字符串。
	err = bucket.PutObject("test.txt", strings.NewReader("hello world!"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func GetOrSave(name string, reader io.Reader) bool {
	isExist, err := bucket.IsObjectExist(name)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !isExist {
		if err := bucket.PutObject(name, reader); err != nil {
		}
	}
	return isExist
}
