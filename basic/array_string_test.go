package basic

import (
	"testing"
)

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
