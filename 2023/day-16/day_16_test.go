package day_16

import (
	_ "embed"
	"reflect"
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
		{"example", args{exampleInput}, 46},
		{"actual", args{actualInput}, 8098},
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
		{"example", args{exampleInput}, 51},
		{"actual", args{actualInput}, 8335},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateActiveFields(t *testing.T) {
	type args struct {
		grid        [][]string
		initialStep NextStep
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateActiveFields(tt.args.grid, tt.args.initialStep); got != tt.want {
				t.Errorf("calculateActiveFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countActiveFields(t *testing.T) {
	type args struct {
		positions []NextStep
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countActiveFields(tt.args.positions); got != tt.want {
				t.Errorf("countActiveFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNextSteps(t *testing.T) {
	type args struct {
		grid      [][]string
		start     [2]int
		direction [2]int
	}
	tests := []struct {
		name         string
		args         args
		wantPossible bool
		wantNext     []NextStep
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPossible, gotNext := getNextSteps(tt.args.grid, tt.args.start, tt.args.direction)
			if gotPossible != tt.wantPossible {
				t.Errorf("getNextSteps() gotPossible = %v, want %v", gotPossible, tt.wantPossible)
			}
			if !reflect.DeepEqual(gotNext, tt.wantNext) {
				t.Errorf("getNextSteps() gotNext = %v, want %v", gotNext, tt.wantNext)
			}
		})
	}
}
