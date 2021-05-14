package hbasic

import (
	"testing"
	"time"
)

func TestTimeFunc_FormatDate(t *testing.T) {
	tests := []struct {
		name string
		t    TimeFunc
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			t: func() time.Time {
				return time.Now()
			},
			want: time.Now().Format("2006-01-02"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.FormatDate(); got != tt.want {
				t.Errorf("FormatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
