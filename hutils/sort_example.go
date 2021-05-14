package hutils

import "fmt"

func exampleStruct() {
	type student struct {
		Name string
		Age  int
	}
	data := []student{
		{Name: "张三", Age: 18},
		{Name: "张四", Age: 32},
		{Name: "张五", Age: 12},
		{Name: "张六", Age: 5},
		{Name: "张七", Age: 43},
		{Name: "张八", Age: 23},
	}
	fmt.Println("before: ", data)
	SortFunc(func() (LenFunc, LessFunc, SwapFunc) {
		return func() int {
				return len(data)
			},
			func(i, j int) bool {
				return data[i].Age < data[j].Age
			},
			func(i, j int) {
				data[i], data[j] = data[j], data[i]
			}
	})
	fmt.Println("after:  ", data)
}
