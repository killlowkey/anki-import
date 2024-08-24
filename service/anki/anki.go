package anki

import (
	"anki-import/anki"
	"anki-import/dict"
	"anki-import/tts"
	"anki-import/util"
	"anki-import/youdao"
	"fmt"
	"log"
)

type AnkiImportService struct {
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

type Option func(ankiImport *AnkiImportService)

// WithSuccessCallback note 导入成功回调
func WithSuccessCallback(success func(word Word, noteId int64)) Option {
	return func(ankiImport *AnkiImportService) {
		ankiImport.success = success
	}
}

// WithFailedCallback note 导入失败回调
func WithFailedCallback(failedCallback func(word Word, err error)) Option {
	return func(ankiImport *AnkiImportService) {
		ankiImport.failed = failedCallback
	}
}

// WithNoteTags deck 笔记的 tags
func WithNoteTags(noteTags []string) Option {
	return func(ankiImport *AnkiImportService) {
		ankiImport.tags = noteTags
	}
}

// WithDebug 开启调试模式
func WithDebug() Option {
	return func(ankiImport *AnkiImportService) {
		ankiImport.ankiClient.SetDebug(true)
	}
}

// WithDict 导入本地字典，csv 格式
func WithDict(wordFilepath, translationFilepath string) Option {
	return func(ankiImport *AnkiImportService) {
		if err := ankiImport.dict.LoadDict(wordFilepath, translationFilepath); err != nil {
			panic(err)
		}
	}
}

func NewImportService(ankiBaseUrl, deckName, modelName string, options ...Option) *AnkiImportService {
	a := &AnkiImportService{
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

func (s *AnkiImportService) ImportNote(importer WordImporter) error {
	if importer == nil {
		return fmt.Errorf("importer is nil")
	}

	words, err := importer.Import()
	if err != nil {
		return err
	}

	return s.ImportNoteWithWords(words)
}

func (s *AnkiImportService) ImportNoteWithWords(words []Word) (err error) {
	var (
		m        map[string]struct{}
		counters int64
	)

	// 从 anki 获取 deck 所有单词
	if _, m, err = s.FindAllWordFromAnki(s.deckName); err != nil {
		return err
	}

	// 对重复单词进行过滤
	words = util.Filter(words, func(word Word) bool {
		if _, ok := m[word.Word]; !ok {
			return true
		} else {
			counters++
			return false
		}
	})

	log.Printf("在【%s】牌组中发现 %d 个重复单词，本次需要导入 %d 个单词\n", s.deckName, counters, len(words))

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

// FindAllWordFromAnki 从 anki 获取名为 deckName 的 deck 所有单词
func (s *AnkiImportService) FindAllWordFromAnki(deckName string) ([]string, map[string]struct{}, error) {
	// 判断是否有该 word，简单处理一下
	noteIds, err := s.ankiClient.FindNotes(fmt.Sprintf("deck:%s", deckName))
	if err != nil {
		return nil, nil, err
	}

	noteInfos, err := s.ankiClient.NotesInfo(noteIds)
	if err != nil {
		return nil, nil, err
	}

	var (
		res []string
		m   = make(map[string]struct{})
	)
	for _, noteInfo := range noteInfos {
		if v, ok := noteInfo.Fields["word"]; ok {
			res = append(res, v.Value)
			m[v.Value] = struct{}{}
		}
	}
	return res, m, nil
}
