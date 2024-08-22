package util

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

func IsWord(text string) bool {
	// 正则表达式：匹配以字母开头，以字母结尾，中间可以有字母、连字符或撇号
	wordRegex := regexp.MustCompile(`^[a-zA-Z]+([-']?[a-zA-Z]+)*$`)
	return wordRegex.MatchString(text)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
