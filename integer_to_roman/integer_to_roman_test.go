package integer_to_roman

import "testing"

func Test_integerToRoman(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerToRoman(tt.args.x); got != tt.want {
				t.Errorf("integerToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
