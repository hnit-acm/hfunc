package hbasic

import (
	"testing"
)

func TestTimeString_GetNative(t *testing.T) {
	tests := []struct {
		name string
		t    TimeString
		want string
	}{
		// TODO: Add test cases.
		{name: "", t: "asdasd", want: "asdasd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.GetNative(); got != tt.want {
				t.Errorf("GetNative() = %v, want %v", got, tt.want)
			}
		})
	}
}
