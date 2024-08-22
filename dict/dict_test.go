package dict

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"testing"
)

func Test_Load(t *testing.T) {
	// 打开 CSV 文件
	file, err := os.OpenFile("word.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	// 设置自定义分隔符，比如使用分号作为分隔符
	gocsv.SetCSVReader(func(reader io.Reader) gocsv.CSVReader {
		r := csv.NewReader(reader)
		r.Comma = '>' // 将分隔符设置为分号
		return r
	})

	var words []Word
	err = gocsv.UnmarshalFile(file, &words)
	if err != nil {
		t.Fatal(err)
	}
}
