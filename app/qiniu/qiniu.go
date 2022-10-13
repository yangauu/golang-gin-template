package qiniu

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"regexp"
	"strconv"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 上传到七牛云 值：目录名、文件
func UploadToQiNiu(path string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 获取凭证、准备上传
	putPlicy := storage.PutPolicy{Scope: Bucket}

	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)

	// 配置信息
	cfg := storage.Config{
		Zone:          &storage.ZoneXinjiapo,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	// 若存在相同文件、则返回上传失败错误
	t := time.Now().Unix()
	s := strconv.FormatInt(t, 10)

	// 处理文件名 原名+时间戳+文件类型
	regfiletype := regexp.MustCompile(`(?:\.\S*)`)
	filetype := regfiletype.FindString(file.Filename)
	regfilename := regexp.MustCompile(`(?:\S*\.)`)
	runefilename := []rune(regfilename.FindString(file.Filename))
	filename := string(runefilename[:len(runefilename)-1])

	log.Println("文件名：", path+filename+s+filetype)
	key := path + filename + s + filetype
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	if err != nil {
		return "", err
	}

	// 拼接路径
	url := ImgUrl + ret.Key
	return url, nil
}

// 删除空间中的文件
func DeleteToQiNiu(filename string) error {
	// 获取凭证、准备上传
	mac := qbox.NewMac(AccessKey, SerectKey)
	cfg := storage.Config{
		UseHTTPS: true, // 是否使用https域名进行资源管理
	}

	// 构建对象
	bucketManager := storage.NewBucketManager(mac, &cfg)

	// 执行删除
	err := bucketManager.Delete(Bucket, filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
