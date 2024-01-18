package day_01

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed example-part-1.txt
var exampleInputPartOne string

//go:embed example-part-2.txt
var exampleInputPartTwo string

func TestPartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{exampleInputPartOne}, 142},
		{"actual", args{input}, 55538},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.input); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{exampleInputPartTwo}, 281},
		{"actual", args{input}, 54875},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
