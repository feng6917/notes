package excelTag

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestWriteXlsx_WriteXlsxFunc(t *testing.T) {
	var ps []Example
	ps = append(ps, Example{
		Id:   111,
		Name: "xxxx",
		Age:  12,
		Sex:  "1",
	})
	ps = append(ps, Example{
		Id:   222,
		Name: "qqqq",
		Age:  22,
		Sex:  "0",
	})
	wx := WriteXlsx{
		SheetName: "People",
		Obj:       ps,
		Funcs: map[string]interface{}{
			"addTen": AddTen,
		},
	}

	err := wx.WriteXlsxFunc()
	if err != nil {
		panic(err)
	}
	f, _ := os.Create("./temp.xlsx")
	_, _ = f.Write(wx.Buffer.Bytes())
}

func AddTen(s string) string {
	i, _ := strconv.Atoi(s)
	return fmt.Sprintf("%d", i+10)
}
