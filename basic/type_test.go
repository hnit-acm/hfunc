package basic

import (
	"reflect"
	"testing"
	"time"
)

func TestJsonString_GetMapStringInterface(t *testing.T) {
	tests := []struct {
		name    string
		j       JsonString
		wantRes map[string]interface{}
	}{
		// TODO: Add test cases.
		{
			name:    "",
			j:       "2020-01-01",
			wantRes: nil,
		},
		{
			name:    "",
			j:       "",
			wantRes: nil,
		},
		{
			name:    "",
			j:       "{\"name\":\"user\"}",
			wantRes: map[string]interface{}{"name": "user"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.j.GetMapStringInterface(); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetMapStringInterface() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestJsonString_GetNative(t *testing.T) {
	tests := []struct {
		name string
		j    JsonString
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			j:    "{}",
			want: "{}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetNative(); got != tt.want {
				t.Errorf("GetNative() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestStringArray_ToString(t *testing.T) {
	type args struct {
		split string
	}
	tests := []struct {
		name    string
		s       ArrayString
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{
			name: "",
			s:    ArrayString{"a", "b", "c"},
			args: args{
				split: "",
			},
			wantRes: "abc",
		},
		{
			name: "",
			s:    ArrayString{},
			args: args{
				split: "",
			},
			wantRes: "",
		},
		{
			name: "",
			s:    ArrayString{"a", "b", "c"},
			args: args{
				split: ",",
			},
			wantRes: "a,b,c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.s.ToString(tt.args.split); gotRes != tt.wantRes {
				t.Errorf("ToString() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func BenchmarkStringArray_ToString(b *testing.B) {
	t := ArrayString{"a", "b", "c"}
	for i := 0; i < b.N; i++ {
		t.ToString("")
	}
}
