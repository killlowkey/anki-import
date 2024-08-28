package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
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

// ExtractTranslations 提取翻译信息并返回 []string 类型结果
func ExtractTranslations(text string) []string {
	// 定义正则表达式模式，匹配类似 "v. " 或 "n. " 等词性标签及其后面的定义
	pattern := `([a-zA-Z]+\.)\s*([^;]+(?:；[^;]+)*)`

	// 编译正则表达式
	regex := regexp.MustCompile(pattern)

	// 查找所有匹配的结果
	matches := regex.FindAllStringSubmatch(text, -1)

	// 创建一个切片来存储词性及其对应的定义
	var results = make([]string, 0)

	// 遍历所有的匹配结果
	for _, match := range matches {
		// match[1] 是词性标签 (e.g., "v." or "n.")
		// match[2] 是对应的定义部分
		partOfSpeech := strings.TrimSpace(match[1])
		definitions := strings.TrimSpace(match[2])
		results = append(results, fmt.Sprintf("%s %s", partOfSpeech, definitions))
	}

	return results
}
