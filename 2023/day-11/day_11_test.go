package day_11

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
		{"example", args{exampleInput}, 374},
		{"actual", args{actualInput}, 9509330},
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
		{"example", args{exampleInput}, 82000210},
		{"actual", args{actualInput}, 635832237682},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateShortestDistances(t *testing.T) {
	type args struct {
		universe        []string
		emptySpaceValue int
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
			if got := calculateShortestDistances(tt.args.universe, tt.args.emptySpaceValue); got != tt.want {
				t.Errorf("calculateShortestDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findExpandingSpaces(t *testing.T) {
	type args struct {
		universe []string
	}
	tests := []struct {
		name           string
		args           args
		wantRowIndices []int
		wantColIndices []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRowIndices, gotColIndices := findExpandingSpaces(tt.args.universe)
			if !reflect.DeepEqual(gotRowIndices, tt.wantRowIndices) {
				t.Errorf("findExpandingSpaces() gotRowIndices = %v, want %v", gotRowIndices, tt.wantRowIndices)
			}
			if !reflect.DeepEqual(gotColIndices, tt.wantColIndices) {
				t.Errorf("findExpandingSpaces() gotColIndices = %v, want %v", gotColIndices, tt.wantColIndices)
			}
		})
	}
}

func Test_findGalaxies(t *testing.T) {
	type args struct {
		universe []string
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGalaxies(tt.args.universe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findGalaxies() = %v, want %v", got, tt.want)
			}
		})
	}
}
