package model

import (
	"context"
	"fmt"
	"gin-gorm/utils"
	"gin-gorm/utils/errmsg"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiNiuServer

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: true,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	fileName := utils.GetCurrentTimeHash()
	err := formUploader.Put(context.Background(), &ret, upToken, fileName, file, fileSize, &putExtra)

	if err != nil {
		fmt.Println(err)
		return "", errmsg.ERROR
	}
	return ImgUrl + "/" + fileName, errmsg.SUCCESS
}


