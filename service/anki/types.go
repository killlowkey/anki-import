package anki

type Word struct {
	Word             string `json:"word" map:"单词"`
	IpaUk            string `json:"ipa_uk" map:"英式音标"`
	IpaUs            string `json:"ipa_us" map:"美式音标"`
	IpaAudio         string `json:"ipa_audio" map:"音频"`
	DefinitionCn     string `json:"definition_cn" map:"翻译"`
	SourceName1      string `json:"source_name1" map:"来源1"`
	SourceContent1   string `json:"source_content1" map:"来源1内容"`
	SourceTranslate1 string `json:"source_translate1" map:"来源1翻译"`
	SourceName2      string `json:"source_name2" map:"来源2"`
	SourceContent2   string `json:"source_content2" map:"来源2内容"`
	SourceTranslate2 string `json:"source_translate2" map:"来源2翻译"`
	Examples1En      string `json:"examples1_en" map:"例子1内容"`
	Examples1Cn      string `json:"examples1_cn" map:"例子1翻译"`
	Examples2En      string `json:"examples2_en" map:"例子2内容"`
	Examples2Cn      string `json:"examples2_cn" map:"例子2翻译"`
}

type WordImporter interface {
	Import() ([]Word, error)
}

func WithXlsx(xlsxFilepath string) WordImporter {
	return WithXlsxAndSheet(xlsxFilepath, "Sheet1")
}

func WithXlsxAndSheet(xlsxFilepath, sheetName string) WordImporter {
	return NewXLSXImporter(xlsxFilepath, sheetName)
}

func WithFeiShuBitTable(appId, appSecret, appToken, tableId string) WordImporter {
	return NewBitTableImporter(appId, appSecret, appToken, tableId)
}
