package office

import "github.com/360EntSecGroup-Skylar/excelize/v2"

type UnitCellIface interface {
	// ColVal 列索引 A-Z
	ColVal() string
	// RowVal 行索引 1——
	RowVal() int
	// Val 单元格值
	Val() interface{}
}

//UnitCell 单位单元格
type UnitCell struct {
	// 列索引 A-Z
	Col string
	// 行索引 1——
	Row int
	// 单元格值
	Value interface{}
}

func (u UnitCell) ColVal() string {
	return u.Col
}

func (u UnitCell) RowVal() int {
	return u.Row
}

func (u UnitCell) Val() interface{} {
	return u.Value
}

func UnitCellIfaceToAxis(iface UnitCellIface) (string, error) {
	return excelize.JoinCellName(iface.ColVal(), iface.RowVal())
}

func AxisToUnitCellIface(axis string) (UnitCellIface, error) {
	col, row, err := excelize.SplitCellName(axis)
	return UnitCell{
		Col: col,
		Row: row,
	}, err
}

func (m UnitCell) GetAxis() (string, error) {
	return excelize.JoinCellName(m.Col, m.Row)
}

func (m UnitCell) CreateColIndex() (next func() UnitCell) {
	index := []rune(m.Col)[0]
	return func() UnitCell {
		index++
		return UnitCell{
			Col: string(index),
			Row: m.Row,
		}
	}
}
