package basic

import (
	"reflect"
	"testing"
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
			if gotRes := tt.j.GetFunc().GetMapStringInterface(); !reflect.DeepEqual(gotRes, tt.wantRes) {
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
