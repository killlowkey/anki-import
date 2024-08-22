package service

import (
	"anki-import/anki"
	"anki-import/dict"
	"anki-import/tts"
	"anki-import/util"
	"anki-import/youdao"
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

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

type AnkiImport struct {
	youDaoClient *youdao.Client
	ankiClient   *anki.Client
	ttsClient    *tts.Client
	dict         *dict.Dict
	deckName     string                        // 添加到那个 deck
	modelName    string                        // note 模板
	tags         []string                      // note 的 tags
	success      func(word Word, noteId int64) // 成功回调
	failed       func(word Word, err error)    // 失败回调
}

type Option func(ankiImport *AnkiImport)

func WithSuccessCallback(success func(word Word, noteId int64)) Option {
	return func(ankiImport *AnkiImport) {
		ankiImport.success = success
	}
}

func WithFailedCallback(failedCallback func(word Word, err error)) Option {
	return func(ankiImport *AnkiImport) {
		ankiImport.failed = failedCallback
	}
}

func WithNoteTags(noteTags []string) Option {
	return func(ankiImport *AnkiImport) {
		ankiImport.tags = noteTags
	}
}

func WithDebug() Option {
	return func(ankiImport *AnkiImport) {
		ankiImport.ankiClient.SetDebug(true)
	}
}

func WithDict(wordFilepath, translationFilepath string, comma rune) Option {
	return func(ankiImport *AnkiImport) {
		// 设置分隔符
		ankiImport.dict.SetComma(comma)
		if err := ankiImport.dict.LoadDict(wordFilepath, translationFilepath); err != nil {
			panic(err)
		}
	}
}

func NewImport(ankiBaseUrl, deckName, modelName string, options ...Option) *AnkiImport {
	a := &AnkiImport{
		youDaoClient: youdao.NewClient(),
		ankiClient:   anki.NewClient(ankiBaseUrl),
		dict:         dict.NewDict(),
		deckName:     deckName,
		modelName:    modelName,
	}

	for _, option := range options {
		option(a)
	}

	return a
}

func (s *AnkiImport) Import(xlsxFilepath string) error {
	return s.ImportWithSheet(xlsxFilepath, "")
}

func (s *AnkiImport) ImportWithSheet(xlsxFilepath, sheetName string) error {
	words, err := s.read(xlsxFilepath, sheetName)
	if err != nil {
		return err
	}

	// 翻译单词含义
	for index, word := range words {
		// 是单词
		if util.IsWord(word.Word) {
			// 翻译
			if word.DefinitionCn == "" {
				if explain, _, err1 := s.youDaoClient.TranslateWord(word.Word); err1 == nil {
					words[index].DefinitionCn = explain
				}
			}

			// 需要添加美式音标
			if w, ok := s.dict.Explain(word.Word); ok {
				words[index].IpaUs = w.PhoneticUS
			}
		}

		// 美英
		words[index].IpaAudio = s.youDaoClient.AudioUS(word.Word)
	}

	// 同步数据
	for _, word := range words {
		note := anki.Note{
			DeckName:  s.deckName,
			ModelName: s.modelName,
			Tags:      s.tags,
		}

		// 处理音频
		if word.IpaAudio != "" {
			note.Audio = []anki.Media{
				{
					URL:      word.IpaAudio,
					Filename: fmt.Sprintf("%s.mp3", word.Word),
					SkipHash: util.GetMD5Hash(word.IpaAudio),
					Fields:   []string{"ipa_audio"}, // 关联音频
				},
			}
			// 需要置空，不然声音那边会显示这个链接
			word.IpaAudio = ""
		}

		// 最后赋值
		note.Fields = word

		// 调用回调
		if noteId, err1 := s.ankiClient.AddNote(note); err1 != nil {
			if s.failed != nil {
				s.failed(word, err1)
			}
		} else {
			if s.success != nil {
				s.success(word, noteId)
			}
		}
	}
	return nil
}

func (s *AnkiImport) read(xlsxFilepath, sheetName string) ([]Word, error) {
	// open an existing file
	wb, err := xlsx.OpenFile(xlsxFilepath)
	if err != nil {
		return nil, err
	}

	// 默认 sheet 名称
	if sheetName == "" {
		sheetName = "Sheet1"
	}

	sh := wb.Sheet[sheetName]

	var data []Word
	for i := 1; i < sh.MaxRow; i++ {
		if sh.Cell(i, 1).String() == "" || sh.MaxCol < 13 {
			continue
		}

		data = append(data, Word{
			Word:             s.getSheetValue(sh, i, 0),
			DefinitionCn:     s.getSheetValue(sh, i, 1),
			IpaUk:            s.getSheetValue(sh, i, 2),
			IpaUs:            s.getSheetValue(sh, i, 3),
			SourceName1:      s.getSheetValue(sh, i, 4),
			SourceContent1:   s.getSheetValue(sh, i, 5),
			SourceTranslate1: s.getSheetValue(sh, i, 6),
			SourceName2:      s.getSheetValue(sh, i, 7),
			SourceContent2:   s.getSheetValue(sh, i, 8),
			SourceTranslate2: s.getSheetValue(sh, i, 9),
			Examples1En:      s.getSheetValue(sh, i, 10),
			Examples1Cn:      s.getSheetValue(sh, i, 11),
			Examples2En:      s.getSheetValue(sh, i, 12),
			Examples2Cn:      s.getSheetValue(sh, i, 13),
		})
	}

	return data, nil
}

func (s *AnkiImport) getSheetValue(sh *xlsx.Sheet, row, col int) string {
	if sh == nil {
		return ""
	}

	res := sh.Cell(row, col).String()
	res = strings.TrimSpace(res)
	res = strings.ReplaceAll(res, "\t", "")
	res = strings.ReplaceAll(res, "\n", " ")
	return res
}
