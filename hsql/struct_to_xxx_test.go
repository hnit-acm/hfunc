package hsql

import (
	"fmt"
	"testing"
)

func TestNewGetFields(t *testing.T) {
	a := struct {
		Name string
		Age  string `alias:"ala"`
	}{}
	fmt.Println(DefaultGetFieldsArray(&a, ""))
}

func BenchmarkNew(b *testing.B) {
	a := struct {
		Name string
		Age  string `alias:"ala"`
	}{}
	for i := 0; i < b.N; i++ {
		DefaultGetFieldsArray(&a, "")
	}
}
