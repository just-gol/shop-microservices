package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	// 定义字符集，包含大小写字母和数字
	const charset = "0123456789"
	var sb strings.Builder

	// 使用 NewSource 创建一个新的随机数源
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// 循环生成指定长度的随机字符串
	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(charset)) // 从charset中随机选择一个字符
		sb.WriteByte(charset[randomIndex])  // 将选中的字符添加到字符串构建器
	}

	return sb.String()
}

func GetOrderId() string {
	template := "20060102150405"
	return time.Now().Format(template) + GenerateRandomString(4)
}
