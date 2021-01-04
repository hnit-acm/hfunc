package basic

import (
	"strconv"
	"testing"
)

func BenchmarkFunc(b *testing.B) {
	_, set, _ := NewHashMapFunc(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set(strconv.Itoa(i), i)
	}
}

func BenchmarkMap(b *testing.B) {
	_, set, _ := NewSyncMapFunc(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set(strconv.Itoa(i), i)
	}
}
