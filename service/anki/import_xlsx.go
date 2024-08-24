package anki

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

type XLSXImporter struct {
	xlsxFilepath string
	sheetName    string
}

func NewXLSXImporter(xlsxFilepath, sheetName string) *XLSXImporter {
	return &XLSXImporter{
		xlsxFilepath: xlsxFilepath,
		sheetName:    sheetName,
	}
}

func (x *XLSXImporter) Import() ([]Word, error) {
	// open an existing file
	wb, err := xlsx.OpenFile(x.xlsxFilepath)
	if err != nil {
		return nil, err
	}

	// 默认 sheet 名称
	if x.sheetName == "" {
		x.sheetName = "Sheet1"
	}

	sh, ok := wb.Sheet[x.sheetName]
	if !ok {
		return nil, fmt.Errorf("not found %s sheet data", x.sheetName)
	}

	var data []Word
	for i := 1; i < sh.MaxRow; i++ {
		if sh.Cell(i, 0).String() == "" || sh.MaxCol < 13 {
			continue
		}

		data = append(data, Word{
			Word:             x.getSheetValue(sh, i, 0),
			DefinitionCn:     x.getSheetValue(sh, i, 1),
			IpaUk:            x.getSheetValue(sh, i, 2),
			IpaUs:            x.getSheetValue(sh, i, 3),
			SourceName1:      x.getSheetValue(sh, i, 4),
			SourceContent1:   x.getSheetValue(sh, i, 5),
			SourceTranslate1: x.getSheetValue(sh, i, 6),
			SourceName2:      x.getSheetValue(sh, i, 7),
			SourceContent2:   x.getSheetValue(sh, i, 8),
			SourceTranslate2: x.getSheetValue(sh, i, 9),
			Examples1En:      x.getSheetValue(sh, i, 10),
			Examples1Cn:      x.getSheetValue(sh, i, 11),
			Examples2En:      x.getSheetValue(sh, i, 12),
			Examples2Cn:      x.getSheetValue(sh, i, 13),
		})
	}

	return data, nil
}

func (x *XLSXImporter) getSheetValue(sh *xlsx.Sheet, row, col int) string {
	if sh == nil {
		return ""
	}

	res := sh.Cell(row, col).String()
	res = strings.TrimSpace(res)
	res = strings.ReplaceAll(res, "\t", "")
	res = strings.ReplaceAll(res, "\n", " ")
	return res
}
