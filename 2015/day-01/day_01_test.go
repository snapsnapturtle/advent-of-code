package day_01

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var actualInput string

func TestPartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"(())"}, 0},
		{"example 2", args{"()()"}, 0},
		{"example 3", args{"((("}, 3},
		{"example 4", args{"(()(()("}, 3},
		{"example 5", args{"))((((("}, 3},
		{"example 6", args{"())"}, -1},
		{"example 7", args{"))("}, -1},
		{"example 8", args{")))"}, -3},
		{"example 9", args{")())())"}, -3},
		{"actual", args{actualInput}, 280},
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
		{"example 1", args{")"}, 1},
		{"example 2", args{"()())"}, 5},
		{"actual", args{actualInput}, 1797},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
