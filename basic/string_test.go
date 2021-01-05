package basic

import "testing"

func TestString_SnakeCasedString(t *testing.T) {
	tests := []struct {
		name string
		s    String
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetFunc().SnakeCasedString(); got != tt.want {
				t.Errorf("SnakeCasedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

//BenchmarkString_SnakeCasedString
//BenchmarkString_SnakeCasedString-8   	 2673525	       437 ns/op
//BenchmarkString_SnakeCasedString-8   	 3009828	       398 ns/op
//BenchmarkString_SnakeCasedString-8   	 3099481	       379 ns/op
//BenchmarkString_SnakeCasedString-8   	 3215841	       371 ns/op
//BenchmarkString_SnakeCasedString-8   	 3209624	       377 ns/op
//BenchmarkString_SnakeCasedString-8   	 3137505	       372 ns/op
//BenchmarkString_SnakeCasedString-8   	 3170169	       381 ns/op
//BenchmarkString_SnakeCasedString-8   	 2699834	       382 ns/op
//BenchmarkString_SnakeCasedString-8   	 3149977	       388 ns/op
//BenchmarkString_SnakeCasedString-8   	 3166250	       400 ns/op
//PASS

func BenchmarkString_SnakeCasedString(b *testing.B) {
	s := "AdasDaDSadADS"
	for i := 0; i < b.N; i++ {
		String(s).GetFunc().SnakeCasedString()
	}
}
