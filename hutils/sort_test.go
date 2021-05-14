package hutils

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func generateTestData(data []int) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		data[i] = random.Intn(10000)
	}
}

func TestSortFunc(t *testing.T) {
	var data = make([]int, 10000)
	generateTestData(data)
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
			SortFunc(tt.args.sortFuncs)
			//SortInts(data)
			fmt.Println(data)
		})
	}
}

func BenchmarkSort(b *testing.B) {
	var data = make([]int, 10000)
	generateTestData(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkSortSlice(b *testing.B) {
	var data = make([]int, 10000)
	generateTestData(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	}
}

func BenchmarkSortFunc(b *testing.B) {
	var data = make([]int, 10000)
	generateTestData(data)
	sortFunc := func() (LenFunc, LessFunc, SwapFunc) {
		return func() int {
				return len(data)
			}, func(i, j int) bool {
				return data[i] > data[j]
			}, func(i, j int) {
				data[i], data[j] = data[j], data[i]
			}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortFunc(sortFunc)
	}
}

func BenchmarkSortInts(b *testing.B) {
	var data = make([]int, 10000)
	generateTestData(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortInts(data)
	}
}
