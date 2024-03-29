package day_02

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var actualInput string

//go:embed example.txt
var exampleInput string

func TestPartOne(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{exampleInput}, 8},
		{"actual", args{actualInput}, 2617},
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
		{"example", args{exampleInput}, 2256},
		{"actual", args{actualInput}, 58604},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
