package anki

import (
	"anki-import/feishu"
	"anki-import/util"
	"context"
)

type BitTableImporter struct {
	client *feishu.BitTableClient
}

func NewBitTableImporter(appId, appSecret, appToken, tableId string) *BitTableImporter {
	return &BitTableImporter{
		client: feishu.NewBitTableClient(appId, appSecret, appToken, tableId),
	}
}

func (b *BitTableImporter) Import() ([]Word, error) {
	recordRespData, err := b.client.List(context.TODO())
	if err != nil {
		return nil, err
	}

	var words []Word
	for _, item := range recordRespData.Items {
		var word Word
		if err1 := util.MapToStruct(item.Fields, &word); err1 == nil {
			words = append(words, word)
		}
	}

	return words, nil
}
