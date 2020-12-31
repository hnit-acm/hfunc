package basic

import (
	"testing"
	"time"
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

func TestTimeString_GetTime(t *testing.T) {
	tests := []struct {
		name    string
		t       TimeString
		wantRes *time.Time
	}{
		// TODO: Add test cases.
		{
			name:    "",
			t:       "2020-01-01",
			wantRes: nil,
		},
		{
			name:    "",
			t:       "",
			wantRes: &time.Time{},
		},
		{
			name:    "",
			t:       "2020-01-01 00:01:01",
			wantRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.t.GetTime(); gotRes == tt.wantRes {
				t.Errorf("GetTime() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func BenchmarkTimeString_GetTime(b *testing.B) {
	t := TimeString("2020-01-")
	t1 := TimeString("2020-01-01 02:21:")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch i % 2 {
		case 0:
			t.GetTime()
		case 1:
			t1.GetTime()
		}
	}
}
