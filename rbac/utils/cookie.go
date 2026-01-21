package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// 定义结构体  缓存结构体 私有
type ginCookie struct{}

// 写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {

	bytes, _ := json.Marshal(value)

	// 加密
	desKey := []byte("itabcd.c")
	encrypt, _ := DesEncrypt(bytes, desKey)
	c.SetCookie(key, string(encrypt), 3600, "/", c.Request.Host, false, true)
}

// 获取数据的方法
func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {

	valueStr, err1 := c.Cookie(key)

	if err1 == nil && valueStr != "" && valueStr != "[]" {
		// 解密
		desKey := []byte("itabcd.c")
		decrypt, e := DesDecrypt([]byte(valueStr), desKey)
		if e != nil {
			return false
		}
		err2 := json.Unmarshal([]byte(decrypt), obj)
		return err2 == nil
	}
	return false
}

func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}

// 实例化结构体
var Cookie = &ginCookie{}
