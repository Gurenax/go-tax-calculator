package main

import "testing"

func Test_convertStringToFloat64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Convert 1.0 to Float", args{"1.0"}, 1.0},
		{"Convert 2.10 to Float", args{"2.10"}, 2.1},
		{"Convert 3.14159 to Float", args{"3.14159"}, 3.14159},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStringToFloat64(tt.args.str); got != tt.want {
				t.Errorf("convertStringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
