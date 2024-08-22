package main

import (
	"anki-import/service"
	"log"
)

func main() {
	ankiImport := service.NewImport(
		"http://localhost:8765",
		"哈利波特与魔法石",
		"english-word",
		service.WithNoteTags([]string{"哈利波特与魔法石"}),
		//service.WithDebug(),
		service.WithSuccessCallback(func(word service.Word, noteId int64) {
			log.Println("添加成功：", noteId)
		}),
		service.WithFailedCallback(func(word service.Word, err error) {
			log.Println("添加失败：", err)
		}),
	)

	err := ankiImport.Import("./testData/哈利波特与魔法石.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
}
