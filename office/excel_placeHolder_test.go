package office

import "testing"

func TestRmPlaceholderSignal(t *testing.T) {
	type args struct {
		signal string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				signal: "{{nieaowei}}",
			},
			want: "nieaowei",
		},
		{
			name: "",
			args: args{
				signal: "{{.nieaowei}}",
			},
			want: "nieaowei",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RmPlaceholderSignal(tt.args.signal); got != tt.want {
				t.Errorf("RmPlaceholderSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSignal(t *testing.T) {
	type args struct {
		signal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				"{{nieaowei}}",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				"12{{nieaowei}}",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSignal(tt.args.signal); got != tt.want {
				t.Errorf("IsSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}
