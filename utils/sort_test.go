package utils

import (
	"fmt"
	"math/rand"
	"testing"
)

var data = make([]int, 10000)

func generateTestData() {
	for i := 0; i < 10000; i++ {
		data[i] = rand.Intn(10000)
	}
}

func TestSortFunc(t *testing.T) {
	type args struct {
		sortFuncs SortFuncs
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				func() (LenFunc, LessFunc, SwapFunc) {
					return func() int {
							return len(data)
						}, func(i, j int) bool {
							return data[i] > data[j]
						}, func(i, j int) {
							data[i], data[j] = data[j], data[i]
						}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateTestData()
			SortFunc(tt.args.sortFuncs)
			fmt.Println(data)
		})
	}
}
