package utils

import (
	"html/template"
	"strconv"
)

// StrToInt String 转 int（十进制）
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StrToIntDefault String 转 int，出错返回 0（看业务情况慎用）
func StrToIntDefault(s string, def int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return i
}

// IntToStr int 转 String
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Float(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	return f, err
}

// 把字符串解析成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

// Substr截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])

}
