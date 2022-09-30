package qiniu

import (
	"context"
	"mime/multipart"

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
	key := path + file.Filename
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	if err != nil {
		return "", err
	}

	// 拼接路径
	url := ImgUrl + ret.Key
	return url, nil
}
