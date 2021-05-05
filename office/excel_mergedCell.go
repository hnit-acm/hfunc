package office

import "github.com/360EntSecGroup-Skylar/excelize/v2"

type MergedCellIface interface {
	UnitCellIface
	StartVal() UnitCellIface
	EndVal() UnitCellIface
	Val() interface{}
}

func IsMergedCell(cell UnitCellIface) (MergedCellIface, bool) {
	val, ok := cell.(MergedCellIface)
	return val, ok
}

type MergedCell struct {
	Start UnitCellIface
	End   UnitCellIface
	Value interface{}
}

func AxisToUnitCell(axis string) (UnitCell, error) {
	col, row, err := excelize.SplitCellName(axis)
	return UnitCell{
		Col: col,
		Row: row,
	}, err
}

func (m MergedCell) StartVal() UnitCellIface {
	return m.Start
}

func (m MergedCell) EndVal() UnitCellIface {
	return m.End
}

func (m MergedCell) Val() interface{} {
	return m.Value
}

func (m MergedCell) ColVal() string {
	return m.StartVal().ColVal()
}

func (m MergedCell) RowVal() int {
	return m.StartVal().RowVal()
}

//MergedIncludeUnit 该合并单元格是否包含某单元格
func MergedIncludeUnit(mergedCell MergedCellIface, cell UnitCellIface) bool {
	if cell.ColVal() >= mergedCell.StartVal().ColVal() && cell.ColVal() <= mergedCell.EndVal().ColVal() &&
		cell.RowVal() >= mergedCell.StartVal().RowVal() && cell.RowVal() <= mergedCell.EndVal().RowVal() {
		return true
	}
	return false
}
