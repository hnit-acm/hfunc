package main

import (
	"fmt"
	"github.com/hnit-acm/hfunc/office"
)

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
	fmt.Println(sheet1.Search(office.SeatSignal, true))
	fmt.Println(excel.SearchSheet("Sheet1", office.SeatSignal, true))
	fmt.Println(sheet1.ParseN())
	excel.SaveAs("./result.xlsx")
	var e office.UnitCellIface = office.UnitCell{}
	_, ok := e.(office.UnitCell)
	fmt.Println(ok)
}
