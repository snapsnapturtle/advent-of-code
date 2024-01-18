package util

import (
	"reflect"
	"testing"
)

func TestParseLinesFromInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty string", args{""}, []string{}},
		{"one line", args{"a"}, []string{"a"}},
		{"two lines", args{"a\nb"}, []string{"a", "b"}},
		{"with terminating new line", args{"a\nb\n"}, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLinesFromInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLinesFromInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
