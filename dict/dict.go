package dict

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	"os"
)

type Dict struct {
	wordMap     map[string]Word
	wordTranMap map[string]Translation
}

type Word struct {
	ID              string  `csv:"vc_id"`               // 单词id
	Vocabulary      string  `csv:"vc_vocabulary"`       // 单词
	PhoneticUK      string  `csv:"vc_phonetic_uk"`      // uk英音音标
	PhoneticUS      string  `csv:"vc_phonetic_us"`      // us美音音标
	Frequency       float64 `csv:"vc_frequency"`        // 词频	0.000000
	Difficulty      int     `csv:"vc_difficulty"`       // 难度	1
	AcknowledgeRate float64 `csv:"vc_acknowledge_rate"` // 认识率	0.664122
	Translation     string  // 从翻译获取
}

type Translation struct {
	Word        string `csv:"word"`        // 单词
	Translation string `csv:"translation"` // 单词的中文翻译
}

func NewDict() *Dict {
	return &Dict{
		wordMap:     make(map[string]Word),
		wordTranMap: make(map[string]Translation),
	}
}

// SetComma 设置分割符
func (d *Dict) SetComma(comma rune) {
	gocsv.SetCSVReader(func(reader io.Reader) gocsv.CSVReader {
		r := csv.NewReader(reader)
		r.Comma = comma
		return r
	})
}

func (d *Dict) LoadDict(wordFilepath, translationFilepath string) error {
	var (
		words        []Word
		translations []Translation
	)

	// 先写死
	d.SetComma('>')

	if err := d.unmarshalFile(wordFilepath, &words); err != nil {
		return err
	} else {
		for _, word := range words {
			d.wordMap[word.Vocabulary] = word
		}
	}

	d.SetComma(',')

	if err := d.unmarshalFile(translationFilepath, &translations); err != nil {
		return err
	} else {
		for _, trans := range translations {
			d.wordTranMap[trans.Word] = trans
		}
	}

	// 组装翻译
	for key, value := range d.wordMap {
		if tran, ok := d.wordTranMap[key]; ok {
			value.Translation = tran.Translation
			d.wordMap[key] = value
		}
	}

	return nil
}

func (d *Dict) unmarshalFile(filepath string, body any) error {
	// 打开 CSV 文件
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将 CSV 数据加载到结构体切片中
	return gocsv.UnmarshalFile(file, body)
}

func (d *Dict) Explain(word string) (Word, bool) {
	if w, ok := d.wordMap[word]; !ok {
		return Word{}, false
	} else {
		return w, true
	}
}
