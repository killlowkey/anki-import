package main

import (
	"anki-import/anki"
	"anki-import/youdao"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/tealeg/xlsx"
	"regexp"
	"strings"
)

var y = youdao.NewClient()

type Word struct {
	Word             string `json:"word" title:"单词"`
	IpaUk            string `json:"ipa_uk" title:"英式音标"`
	IpaUs            string `json:"ipa_us" title:"美式音标"`
	IpaAudio         string `json:"ipa_audio" title:"音频"`
	DefinitionCn     string `json:"definition_cn" title:"翻译"`
	SourceName1      string `json:"source_name1" title:"来源1"`
	SourceContent1   string `json:"source_content1" title:"来源1内容"`
	SourceTranslate1 string `json:"source_translate1" title:"来源1翻译"`
	SourceName2      string `json:"source_name2" title:"来源2"`
	SourceContent2   string `json:"source_content2" title:"来源2内容"`
	SourceTranslate2 string `json:"source_translate2" title:"来源2翻译"`
	Examples1En      string `json:"examples1_en" title:"例子1内容"`
	Examples1Cn      string `json:"examples1_cn" title:"例子1翻译"`
	Examples2En      string `json:"examples2_en" title:"例子2内容"`
	Examples2Cn      string `json:"examples2_cn" title:"例子2翻译"`
}

func isWord(text string) bool {
	// 正则表达式：匹配以字母开头，以字母结尾，中间可以有字母、连字符或撇号
	wordRegex := regexp.MustCompile(`^[a-zA-Z]+([-']?[a-zA-Z]+)*$`)
	return wordRegex.MatchString(text)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getSheetValue(sh *xlsx.Sheet, row, col int) string {
	if sh == nil {
		return ""
	}

	res := sh.Cell(row, col).String()
	res = strings.TrimSpace(res)
	res = strings.ReplaceAll(res, "\t", "")
	res = strings.ReplaceAll(res, "\n", " ")
	return res
}

func read(file string) ([]Word, error) {
	// open an existing file
	wb, err := xlsx.OpenFile("./testData/哈利波特与魔法石.xlsx")
	if err != nil {
		return nil, err
	}

	sh := wb.Sheet["Sheet1"]

	var data []Word
	for i := 1; i < sh.MaxRow; i++ {
		if sh.Cell(i, 1).String() == "" || sh.MaxCol < 13 {
			continue
		}

		data = append(data, Word{
			Word:             getSheetValue(sh, i, 0),
			IpaUk:            getSheetValue(sh, i, 1),
			IpaUs:            getSheetValue(sh, i, 2),
			DefinitionCn:     getSheetValue(sh, i, 3),
			SourceName1:      getSheetValue(sh, i, 4),
			SourceContent1:   getSheetValue(sh, i, 5),
			SourceTranslate1: getSheetValue(sh, i, 6),
			SourceName2:      getSheetValue(sh, i, 7),
			SourceContent2:   getSheetValue(sh, i, 8),
			SourceTranslate2: getSheetValue(sh, i, 9),
			Examples1En:      getSheetValue(sh, i, 10),
			Examples1Cn:      getSheetValue(sh, i, 11),
			Examples2En:      getSheetValue(sh, i, 12),
			Examples2Cn:      getSheetValue(sh, i, 13),
		})
	}

	return data, nil
}

func main() {
	words, err := read("./testData/哈利波特与魔法石.xlsx")
	if err != nil {
		panic(err)
	}
	// 处理单词
	for index, word := range words {
		// 不是单词
		if !isWord(word.Word) {
			continue
		}

		if word.DefinitionCn == "" {
			if explain, _, err1 := y.TranslateWord(word.Word); err1 == nil {
				words[index].DefinitionCn = explain
			}
		}

		words[index].IpaAudio = y.AudioUS(word.Word)
	}

	// 同步数据
	for _, word := range words {
		note := anki.Note{
			DeckName:  "哈利波特与魔法石",
			ModelName: "english-word",
			Tags:      []string{"哈利波特与魔法石"},
		}

		if word.IpaAudio != "" {
			note.Audio = []anki.Media{
				{
					URL:      word.IpaAudio,
					Filename: fmt.Sprintf("%s.mp3", word.Word),
					SkipHash: getMD5Hash(word.IpaAudio),
					Fields:   []string{"ipa_audio"}, // 关联音频
				},
			}
			// 需要置空，不然声音那边会显示这个链接
			word.IpaAudio = ""
		}

		// 最后赋值
		note.Fields = word

		client := anki.NewClient("http://localhost:8765", anki.WithDebug())
		if _, err2 := client.AddNote(note); err2 != nil {
			panic(err2)
		}
	}
}
