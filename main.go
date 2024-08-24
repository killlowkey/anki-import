package main

import (
	"anki-import/service/anki"
	"log"
	"os"
)

func main() {
	ankiImport := anki.NewImportService(
		"http://localhost:8765",
		"哈利波特与魔法石",
		"english-word",
		anki.WithNoteTags([]string{"哈利波特与魔法石"}),
		//service.WithDebug(),
		anki.WithDict("./dict/word.csv", "./dict/word_translation.csv"),
		anki.WithSuccessCallback(func(word anki.Word, noteId int64) {
			log.Println("添加成功：", noteId)
		}),
		anki.WithFailedCallback(func(word anki.Word, err error) {
			log.Println("添加失败：", err)
		}),
	)

	// anki.WithXlsx("./testData/哈利波特与魔法石.xlsx")
	err := ankiImport.ImportNote(anki.WithFeiShuBitTable(
		os.Getenv("FEISHU_APP_ID"),
		os.Getenv("FEISHU_APP_SECRET"),
		os.Getenv("FEISHU_APP_TOKEN"),
		os.Getenv("FEISHU_APP_TABLE_ID"),
	))
	if err != nil {
		log.Fatalln(err)
	}
}
