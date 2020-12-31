package utils

import "github.com/hnit-acm/go-common/basic"

func SnakeCasedString(str string) string {
	return basic.String(str).SnakeCasedString()
}
