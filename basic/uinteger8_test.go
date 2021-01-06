package basic

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	number := UInteger8(8)
	fmt.Println(number)
	fmt.Println(number.GetFunc()())
	fmt.Println(number.GetNative())
}


func Test02(t *testing.T) {
	number := UInteger8(8)
	numberFunc := number.GetFunc()
	numberFunc.GetNative()
}

func TestUInteger8Func_ToHexStringUpper(t *testing.T) {
	fmt.Println(UInteger8(123).GetFunc().ToHexStringUpper())
	fmt.Println(UInteger8(123).GetFunc().ToHexStringLower())
}