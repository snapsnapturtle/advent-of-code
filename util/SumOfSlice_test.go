package util

import "testing"

func TestSumOfSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty slice", args{[]int{}}, 0},
		{"one element", args{[]int{1}}, 1},
		{"three elements", args{[]int{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfSlice(tt.args.slice); got != tt.want {
				t.Errorf("SumOfSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
