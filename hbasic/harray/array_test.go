package harray

import (
	"fmt"
	"testing"
)

var array []string

func init() {
	array = make([]string, 0)
	for i := 10000000; i < 10020000; i++ {
		array = append(array, fmt.Sprintf("%d", i))
	}
}

func BenchmarkArrayStringToString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToString(array, ",")
	}
}
