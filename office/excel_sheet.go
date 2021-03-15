package office

import (
	"errors"
	"fmt"
	"github.com/hnit-acm/hfunc/utils"
	"strings"
)

// SheetFunc 工作表操作函数
type SheetFunc struct {
	e *Excel
}

func (e *SheetFunc) UnmergeCell(mergeCell MergedCellIface) error {
	h, err := UnitCellIfaceToAxis(mergeCell.StartVal())
	if err != nil {
		return err
	}
	v, err := UnitCellIfaceToAxis(mergeCell.EndVal())
	if err != nil {
		return err
	}
	return e.e.UnmergeCell(e.e.currentSheet, h, v)
}

// CreateColIndexWithMerged 创建列索引函数，以合并单元格的列为准
func (e *SheetFunc) CreateColIndexWithMerged(start UnitCellIface) func() UnitCellIface {
	index := []rune(start.ColVal())[0] - 1
	tempCell := UnitCell{
		Col: string(index),
		Row: start.RowVal(),
	}
	mergeCells, _ := e.GetMergedCells()
	isInclude := func(cell UnitCell) (MergedCellIface, bool) {
		for _, mergeCell := range mergeCells {
			if MergedIncludeUnit(mergeCell, cell) {
				return mergeCell, true
			}
		}
		return nil, false
	}

	return func() UnitCellIface {
		// 如果是合并单元格
		if cell, ok := isInclude(tempCell); ok {
			index = []rune(cell.EndVal().ColVal())[0] + 1
		} else {
			index++
		}
		tempCell.Col = string(index)
		fmt.Println(tempCell)
		return tempCell
	}
}

// CreateRowIndexWithMerged 创建行索引函数，以合并单元格的列为准
func (e *SheetFunc) CreateRowIndexWithMerged(start UnitCellIface) func() UnitCellIface {
	index := start.RowVal() - 1
	tempCell := UnitCell{
		Col: start.ColVal(),
		Row: index,
	}
	mergeCells, _ := e.GetMergedCells()
	isInclude := func(cell UnitCell) (MergedCellIface, bool) {
		for _, mergeCell := range mergeCells {
			if MergedIncludeUnit(mergeCell, cell) {
				return mergeCell, true
			}
		}
		return nil, false
	}
	return func() UnitCellIface {
		// 如果是合并单元格
		if cell, ok := isInclude(tempCell); ok {
			index = cell.EndVal().RowVal() + 1
		} else {
			index++
		}
		tempCell.Row = index
		fmt.Println(tempCell)
		return tempCell
	}
}

// SetRowWithMerged 按单位单元格或合并单元格（取决于excel中格式）设置值
func (e *SheetFunc) SetRowWithMerged(startAxis UnitCellIface, slice interface{}) error {
	nextColCell := e.CreateColIndexWithMerged(startAxis)
	var data []interface{}
	err := utils.SourceToTarget(slice, &data)
	if err != nil {
		return err
	}
	for _, datum := range data {
		err := e.SetCellValue(nextColCell(), datum)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetColWithMerged 按单位单元格或合并单元格（取决于excel中格式）设置值
func (e *SheetFunc) SetColWithMerged(startAxis UnitCellIface, slice interface{}) error {
	nextRowCell := e.CreateRowIndexWithMerged(startAxis)
	data := slice.([]interface{})
	for _, datum := range data {
		err := e.SetCellValue(nextRowCell(), datum)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetRow 强制按单元格设置值
func (e *SheetFunc) SetRow(startAxis UnitCellIface, slice interface{}) error {
	a, err := UnitCellIfaceToAxis(startAxis)
	if err != nil {
		return err
	}
	return e.e.SetSheetRow(e.e.currentSheet, a, slice)
}

// SetMergeCellValue 设置合并单元格值，本质上市调用 SetCellValue，因为设置合并单元格的值
// 就是设置合并单元格左上角的单位单元格的值
func (e *SheetFunc) SetMergeCellValue(axis MergedCellIface, value interface{}) error {
	return e.SetCellValue(axis, value)
}

// SetCellValue 设置单位单元格的值
func (e *SheetFunc) SetCellValue(axis UnitCellIface, value interface{}) error {
	a, err := UnitCellIfaceToAxis(axis)
	if err != nil {
		return err
	}
	return e.e.SetCellValue(e.e.currentSheet, a, value)
}

// GetMergedCells 获取所有合并单元格的值
func (e *SheetFunc) GetMergedCells() (res []MergedCellIface, err error) {
	mergeCells, err := e.e.GetMergeCells(e.e.currentSheet)
	if err != nil {
		return
	}
	for _, cell := range mergeCells {
		axises := strings.Split(cell[0], ":")
		if len(axises) != 2 {
			return nil, errors.New("axis format error")
		}
		startAxis, err := AxisToUnitCell(axises[0])
		if err != nil {
			return nil, err
		}
		endAxis, err := AxisToUnitCell(axises[1])
		if err != nil {
			return nil, err
		}
		res = append(res, MergedCell{
			Start: startAxis,
			End:   endAxis,
			Value: cell[1],
		})
	}
	return
}

// Search 搜索所有匹配单元格
func (e *SheetFunc) Search(str string, reg bool) (res []UnitCellIface, err error) {
	// 搜索表 拿去所有匹配单元格的位置（包括合并单元格 ）
	unitAxisList, err := e.e.SearchSheet(e.e.currentSheet, str, reg)
	if err != nil {
		return
	}
	// 获取合并单元格，用于合并数据
	mergedCells, err := e.GetMergedCells()
	// 构建包含函数
	includeCell := func(cell UnitCellIface) (MergedCellIface, bool) {
		for _, mergedCell := range mergedCells {
			if MergedIncludeUnit(mergedCell, cell) {
				return mergedCell, true
			}
		}
		return nil, false
	}
	if err != nil {
		return nil, err
	}
	// 合并数据
	for _, unit := range unitAxisList {
		cell, err := AxisToUnitCell(unit)
		if err != nil {
			return nil, err
		}
		// 如果属于合并单元格，则返回合并单元格
		if v, ok := includeCell(cell); ok {
			val, err := e.GetCellValue(cell)
			if err != nil {
				return nil, err
			}
			res = append(res, MergedCell{
				Start: v.StartVal(),
				End:   v.EndVal(),
				Value: val,
			})
		} else { //否则 返回单位单元格
			cell.Value, err = e.GetCellValue(cell)
			if err != nil {
				return nil, err
			}
			res = append(res, cell)
		}
	}
	return
}

// GetCellValue 获取单位单元格的值
func (e *SheetFunc) GetCellValue(cell UnitCellIface) (string, error) {
	axis, err := UnitCellIfaceToAxis(cell)
	if err != nil {
		return "", err
	}
	return e.e.GetCellValue(e.e.currentSheet, axis)
}

// GetMergedCellValue 获取合并单元格的值
func (e *SheetFunc) GetMergedCellValue(cell UnitCellIface) (string, error) {
	return e.GetCellValue(cell)
}

func (e *SheetFunc) SetRowValueStartFrom(startCell UnitCellIface, slice [][]interface{}, length int) error {
	nextCol := e.CreateRowIndexWithMerged(startCell)
	for index, row := range slice {
		if index >= length && length > 0 {
			break
		}
		err := e.SetRowWithMerged(nextCol(), row)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *SheetFunc) SetColValueStartFrom(startCell UnitCellIface, slice [][]interface{}, length int) error {
	nextCol := e.CreateColIndexWithMerged(startCell)
	for index, row := range slice {
		if index >= length {
			break
		}
		err := e.SetRowWithMerged(nextCol(), row)
		if err != nil {
			return err
		}
	}
	return nil
}
