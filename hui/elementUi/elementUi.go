package elementUi

// 选择框映射数据 支持级联
type SelectItem struct {
	Label string      `json:"label"`
	Value interface{} `json:"key"`
}

type SelectData []SelectItem

// 级联选择映射数据
type CascadeItem struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"key"`
	Children CascadeData
}

type CascadeData []CascadeItem

func main() {

}
