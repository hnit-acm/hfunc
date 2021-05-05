package utils

import (
	"fmt"
)

func AnyToString(any interface{}) string {
	return fmt.Sprintf("%v", any)
}
