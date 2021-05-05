package main

import (
	"fmt"
	"github.com/hnit-acm/hfunc/office"
)

type ExcelOption struct {
	Attr  string
	Value []string
}

func (e ExcelOption) NoEmptyHandler(excel *office.SheetFunc, placeholder office.PlaceholderIface, signal string) (string, error) {
	excel.SetRowWithMerged(placeholder.CellVal(), e.Value)
	return "", nil
}

func (e ExcelOption) EmptyHandler(excel *office.SheetFunc, placeholder office.PlaceholderIface, signal string) (string, error) {
	return "空", nil
}

func main() {
	excel, _ := office.OpenExcelFromFile("./test.xlsx")
	sheet1 := excel.SelectSheet("Sheet1")
	sheet1.SetRowWithMerged(office.UnitCell{
		Col: "P",
		Row: 1,
	}, []interface{}{
		"nieaowei", "123", "123wdas", "dsadqwewq", "dioasdoihasd",
	})
	sheet1.SetRowWithMerged(office.UnitCell{
		Col: "P",
		Row: 4,
	}, []interface{}{
		"nieaowei", "123", "123wdas", "dsadqwewq", "dioasdoihasd",
	})
	sheet1.SetRowWithMerged(office.UnitCell{
		Col: "P",
		Row: 7,
	}, []interface{}{
		"nieaowei", "123", "123wdas", "dsadqwewq", "dioasdoihasd",
	})
	sheet1.SetColWithMerged(office.UnitCell{
		Col: "K",
		Row: 1,
	}, []interface{}{
		"nieaowei", "123", "123wdas", "dsadqwewq", "dioasdoihasd",
	})
	sheet1.SetRowValueStartFrom(office.UnitCell{
		Col: "L",
		Row: 1,
	}, [][]interface{}{
		{"123", "123", "123"},
		{"123", "123", "123"},
		{"123", "123", "123"},
		{"123", "123", "123"},
	}, 3)
	parseList, _ := sheet1.Parse()
	for _, iface := range parseList {
		sheet1.SetPlaceholder(iface, map[string]interface{}{
			"nieaowei": "fyw是傻逼",
			"123123": ExcelOption{
				Attr:  "123123",
				Value: []string{"dsadsa", "1231313", "dsadqweqeasd"},
			},
		})
	}
	excel.SaveAs("./result.xlsx")
	var e office.UnitCellIface = office.UnitCell{}
	_, ok := e.(office.UnitCell)
	fmt.Println(ok)
}
