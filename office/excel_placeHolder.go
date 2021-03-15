package office

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"regexp"
	"strings"
)

// {{.name}} 值
// 占位符
// {{name}} 值
// {{name.c}} 按单位列
// {{name.r}} 按单位行
// {{name.mc}} 按合并列
// {{name.mr}} 按合并行

const NameRegStr = `([A-Z]|[a-z]|[0-9]|_|)+`
const ValRegStr = `\.` + NameRegStr

const SeatSignal = `{{([A-Z]|[a-z]|[0-9]|_|\.)+}}`

const SeatValueSignal = `^{{\.([A-Z]|[a-z]|[0-9]|_)+}}$`

var SeatReg, _ = regexp.Compile(SeatSignal)
var SeatValueReg, _ = regexp.Compile(SeatValueSignal)

type PlaceholderHandler interface {
	NoEmptyHandler(excel *SheetFunc, placeholder PlaceholderIface, signal string) (string, error)
	EmptyHandler(excel *SheetFunc, placeholder PlaceholderIface, signal string) (string, error)
}

func IsPlaceholderHandler(data interface{}) (PlaceholderHandler, bool) {
	val, ok := data.(PlaceholderHandler)
	return val, ok
}

type PlaceholderIface interface {
	SignalsVal() []string
	CellVal() UnitCellIface
}

type Placeholder struct {
	Signals []string
	Cell    UnitCellIface
}

func (p Placeholder) SignalsVal() []string {
	return p.Signals
}

func (p Placeholder) CellVal() UnitCellIface {
	return p.Cell
}

func (e *SheetFunc) SetPlaceholder(placeholder PlaceholderIface, data map[string]interface{}) error {
	text := placeholder.CellVal().Val()
	for _, signalVal := range placeholder.SignalsVal() {
		val, ok := data[signalVal]
		// 如果有数据
		if ok {
			if placeholderHandler, ok := IsPlaceholderHandler(val); ok {
				noEmpty, err := placeholderHandler.NoEmptyHandler(e, placeholder, signalVal)
				if err != nil {
					return err
				}
				if noEmpty == "" {
					continue
				}
				val = noEmpty
			}
			strings.ReplaceAll(text.(string), signalVal, val.(string))
		} else { // 如果没有数据
			if placeholderHandler, ok := IsPlaceholderHandler(val); ok {
				noEmpty, err := placeholderHandler.EmptyHandler(e, placeholder, signalVal)
				if err != nil {
					return err
				}
				strings.ReplaceAll(text.(string), signalVal, noEmpty)
				continue
			}
			strings.ReplaceAll(text.(string), signalVal, "")
		}
	}
	if mergedCell, ok := IsMergedCell(placeholder.CellVal()); ok {
		err := e.SetMergeCellValue(mergedCell, text)
		if err != nil {
			return err
		}
	} else {
		err := e.SetCellValue(placeholder.CellVal(), text)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *SheetFunc) ParseN() (res []PlaceholderIface, err error) {
	placeHolders, err := e.Search(SeatSignal, true)
	if err != nil {
		return
	}
	for _, placeholder := range placeHolders {
		item := Placeholder{
			Signals: SeatReg.FindAllString(placeholder.Val().(string), -1),
			Cell:    placeholder,
		}
		res = append(res, item)
	}
	return
}

type AxisItem struct {
	Start   string
	End     string
	IsMerge bool
}

type SeatItem struct {
	Text string
	Name []string
	Axis AxisItem
}

func (a AxisItem) CreateRowIndex() (next func() string) {
	index := a.Start
	return func() string {
		c, r, _ := excelize.SplitCellName(index)
		index, _ = excelize.JoinCellName(c, r+1)
		return index
	}
}
