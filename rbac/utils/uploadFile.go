package utils

import (
	"errors"
	"fmt"
	"gin-shop/models"
	"mime/multipart"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	. "github.com/hunterhug/go_image"
	"gopkg.in/ini.v1"
)

// 上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	ossStatus := GetOssStatus()
	if ossStatus == 1 {
		return OssUploadImg(c, picName)
	} else {
		return LocalUploadImg(c, picName)
	}
}

// 上传图片到Oss
func OssUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、定义图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "static/upload/" + day

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixN(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)

	OssUplod(file, dst)
	return dst, nil

}

// 上传图片到本地
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	fmt.Println(file)
	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、创建图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "./static/upload/" + day

	err1 := os.MkdirAll(dir, 0777)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixN(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil

}

// 获取Oss的状态
func GetOssStatus() int {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus, _ := StrToInt(config.Section("oss").Key("status").String())
	return ossStatus
}

// 通过列获取值
func GetSettingFromColumn(columnName string) string {
	//redis file
	setting := models.Setting{}
	models.DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}
func FormatImg(str string) string {
	ossStatus := GetOssStatus()
	if ossStatus == 1 {
		return GetSettingFromColumn("OssDomain") + str
	} else {
		return "/" + str
	}
}

//生成商品缩略图

func ResizeGoodsImage(filename string) {
	extname := path.Ext(filename)
	ThumbnailSize := strings.ReplaceAll(GetSettingFromColumn("ThumbnailSize"), "，", ",")
	thumbnailSizeSlice := strings.Split(ThumbnailSize, ",")
	//static/upload/tao_400.png
	//static/upload/tao_400.png_100x100.png
	for i := 0; i < len(thumbnailSizeSlice); i++ {
		savePath := filename + "_" + thumbnailSizeSlice[i] + "x" + thumbnailSizeSlice[i] + extname
		w, _ := StrToInt(thumbnailSizeSlice[i])
		err := ThumbnailF2F(filename, savePath, w, w)
		if err != nil {
			Logger.Error("生成商品缩略图失败", err)
		}
	}

}

func FormatAttr(str string) string {

	tempSlice := strings.Split(str, "\n")
	var tempStr string
	for _, v := range tempSlice {
		md := []byte(v)
		output := markdown.ToHTML(md, nil, nil)
		tempStr += string(output)
	}
	return tempStr
}
