package office

import (
	"github.com/hnit-acm/hfunc/utils"
	"regexp"
	"strings"
)

// 占位符
// {{name}} 值
// {{.name}} 值
// {{name.c}} 按单位列 -- 未来
// {{name.r}} 按单位行 -- 未来
// {{name.mc}} 按合并列 -- 未来
// {{name.mr}} 按合并行 -- 未来

const signalNameRegStr = `([A-Z]|[a-z]|[0-9]|_|)+`
const signalValRegStr = `\.` + signalNameRegStr
const SignalRegStr = `{{(` + `(` + signalNameRegStr + `)|(` + signalValRegStr + `)` + `)}}`

var SignalReg, _ = regexp.Compile(SignalRegStr)
var SignalValReg, _ = regexp.Compile(signalValRegStr)
var SignalRegAll, _ = regexp.Compile("^" + SignalRegStr + "$")
var SignalValRegAll, _ = regexp.Compile("^" + signalValRegStr + "$")

func RmPlaceholderSignal(signal string) string {
	return strings.Trim(signal, "({{)|.|(}})")
}

func IsSignal(signal string) bool {
	return SignalRegAll.MatchString(signal)
}

func IsSignalVal(signal string) bool {
	return SignalValRegAll.MatchString(signal)
}

type EmptyHandler interface {
	EmptyHandle(sheet *SheetFunc, placeholder PlaceholderIface, signal string) (string, error)
}

type NoEmptyHandler interface {
	NoEmptyHandle(excel *SheetFunc, placeholder PlaceholderIface, signal string) (string, error)
}

func IsNoEmptyHandler(data interface{}) (NoEmptyHandler, bool) {
	val, ok := data.(NoEmptyHandler)
	return val, ok
}

func IsEmptyHandler(data interface{}) (EmptyHandler, bool) {
	val, ok := data.(EmptyHandler)
	return val, ok
}

type PlaceholderHandler interface {
	NoEmptyHandler
	EmptyHandler
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
	// 处理数据
	text := placeholder.CellVal().Val()
	for _, signalVal := range placeholder.SignalsVal() {
		key := RmPlaceholderSignal(signalVal)
		val, ok := data[key]
		// 如果有数据
		if ok {
			// 如果实现了处理接口
			if placeholderHandler, ok := IsNoEmptyHandler(val); ok {
				noEmpty, err := placeholderHandler.NoEmptyHandle(e, placeholder, signalVal)
				if err != nil {
					return err
				}
				if noEmpty == "" {
					return nil
				}
				val = noEmpty
			}
			text = strings.ReplaceAll(utils.AnyToString(text), signalVal, utils.AnyToString(val))
		} else { // 如果没有数据
			// 如果实现了处理接口
			if placeholderHandler, ok := IsEmptyHandler(val); ok {
				noEmpty, err := placeholderHandler.EmptyHandle(e, placeholder, signalVal)
				if err != nil {
					return err
				}
				// 如果为空，则说明在handler已经处理了数据
				if noEmpty == "" {
					return nil
				}
				strings.ReplaceAll(text.(string), signalVal, noEmpty)
				continue
			}
			text = strings.ReplaceAll(text.(string), signalVal, "")
		}
	}
	// 填充数据
	// 如果是合并单元格
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

func (e *SheetFunc) Parse() (res []PlaceholderIface, err error) {
	placeHolders, err := e.Search(SignalRegStr, true)
	if err != nil {
		return
	}
	for _, placeholder := range placeHolders {
		item := Placeholder{
			Signals: SignalReg.FindAllString(placeholder.Val().(string), -1),
			Cell:    placeholder,
		}
		res = append(res, item)
	}
	return
}
