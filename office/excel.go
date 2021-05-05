package office

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
)

type Excel struct {
	*excelize.File
	currentSheet string
}

// SelectSheet 选择工作表
func (e *Excel) SelectSheet(sheetName ...string) *SheetFunc {
	e.currentSheet = e.GetSheetName(0)
	for _, s := range sheetName {
		e.currentSheet = s
	}
	f := SheetFunc{
		e,
	}
	return &f
}

// OpenExcelFromFile 从文件读取excel
func OpenExcelFromFile(filename string) (*Excel, error) {
	f, err := excelize.OpenFile(filename)
	return &Excel{File: f}, err
}

// OpenExcelFromReader 从读取流读取excel
func OpenExcelFromReader(reader io.Reader) (*Excel, error) {
	f, err := excelize.OpenReader(reader)
	return &Excel{File: f}, err
}
