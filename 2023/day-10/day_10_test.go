package day_10

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var actualInput string

//go:embed example-part-1.txt
var exampleInputPartOne string

//go:embed example-part-2-1.txt
var exampleInputPartTwoA string

//go:embed example-part-2-2.txt
var exampleInputPartTwoB string

//go:embed example-part-2-3.txt
var exampleInputPartTwoC string

func TestPartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{exampleInputPartOne}, 4},
		{"actual", args{actualInput}, 6923},
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
		{"example 1", args{exampleInputPartTwoA}, 4},
		{"example 2", args{exampleInputPartTwoB}, 8},
		{"example 3", args{exampleInputPartTwoC}, 10},
		{"actual", args{actualInput}, 529},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
