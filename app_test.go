package main

import (
	"fmt"
	"reflect"
	"testing"
)

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
		{"Convert a letter", args{"a"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStringToFloat64(tt.args.str); got != tt.want {
				t.Errorf("convertStringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStringToFloat64_DataType(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Convert 1.0 to Float", args{"1.0"}, "float64"},
		{"Convert 2.10 to Float", args{"2.10"}, "float64"},
		{"Convert 3.14159 to Float", args{"3.14159"}, "float64"},
		{"Convert a letter", args{"a"}, "float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStringToFloat64(tt.args.str); fmt.Sprintf("%v", reflect.TypeOf(got)) != tt.want {
				t.Errorf("convertStringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStringToFloat64_DataType_WithoutReflect(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Convert 1.0 to Float", args{"1.0"}, "float64"},
		{"Convert 2.10 to Float", args{"2.10"}, "float64"},
		{"Convert 3.14159 to Float", args{"3.14159"}, "float64"},
		{"Convert a letter", args{"a"}, "float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStringToFloat64(tt.args.str); fmt.Sprintf("%T", got) != tt.want {
				t.Errorf("convertStringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineTaxBracket(t *testing.T) {
	type args struct {
		salary float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Zero", args{0}, "None"},
		{"Less than Zero", args{-1}, "None"},
		{"Less than 18200", args{18199.0}, "Tax Bracket 1"},
		{"Equals 18200", args{18200.0}, "Tax Bracket 1"},
		{"Greater than 18200", args{18201.0}, "Tax Bracket 2"},
		{"Less than 37000", args{36999.0}, "Tax Bracket 2"},
		{"Equals 37000", args{37000.0}, "Tax Bracket 2"},
		{"Greater than 37000", args{37001.0}, "Tax Bracket 3"},
		{"Less than 90000", args{89000.0}, "Tax Bracket 3"},
		{"Equals 90000", args{90000.0}, "Tax Bracket 3"},
		{"Greater than 90000", args{90001.0}, "Tax Bracket 4"},
		{"Less than 180000", args{179000.0}, "Tax Bracket 4"},
		{"Equals 180000", args{180000.0}, "Tax Bracket 4"},
		{"Greater than 180000", args{180001.0}, "Tax Bracket 5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineTaxBracket(tt.args.salary); got != tt.want {
				t.Errorf("determineTaxBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeTax(t *testing.T) {
	type args struct {
		salary float64
		num1   float64
		num2   float64
		num3   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Compute taxced salary for $18201", args{18201.0, 18200.0, 0.19, 0.0}, 0.19},
		{"Compute taxced salary for $20000", args{20000.0, 18200.0, 0.19, 0.0}, 342.00},
		{"Compute taxced salary for $50000", args{50000.0, 37000.0, 0.325, 3572.0}, 7797.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeTax(tt.args.salary, tt.args.num1, tt.args.num2, tt.args.num3); got != tt.want {
				t.Errorf("computeTax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateAnnualTax(t *testing.T) {
	type args struct {
		salary float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Annual tax for $10000", args{10000}, 0},
		{"Annual tax for $20000", args{20000}, 341.81},
		{"Annual tax for $30000", args{30000}, 2241.81},
		{"Annual tax for $40000", args{40000}, 3466.6749999999997},
		{"Annual tax for $50000", args{50000}, 6716.675},
		{"Annual tax for $50000", args{100000}, 22066.63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAnnualTax(tt.args.salary); got != tt.want {
				t.Errorf("calculateAnnualTax() = %v, want %v", got, tt.want)
			}
		})
	}
}
