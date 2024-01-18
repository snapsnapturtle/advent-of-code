package day_19

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
		{"example", args{exampleInput}, 19114},
		{"actual", args{actualInput}, 509597},
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
	t.Skip("Solution implemented")

	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{exampleInput}, 167409079868000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart_matchesCondition(t *testing.T) {
	type fields struct {
		X int
		M int
		A int
		S int
	}
	type args struct {
		field     string
		condition string
		value     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Part{
				X: tt.fields.X,
				M: tt.fields.M,
				A: tt.fields.A,
				S: tt.fields.S,
			}
			if got := p.matchesCondition(tt.args.field, tt.args.condition, tt.args.value); got != tt.want {
				t.Errorf("matchesCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflowStep_isNextStepFinal(t *testing.T) {
	type fields struct {
		Field          string
		Condition      string
		Value          int
		NextWorkflowId string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WorkflowStep{
				Field:          tt.fields.Field,
				Condition:      tt.fields.Condition,
				Value:          tt.fields.Value,
				NextWorkflowId: tt.fields.NextWorkflowId,
			}
			if got := w.isNextStepFinal(); got != tt.want {
				t.Errorf("isNextStepFinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPartAccepted(t *testing.T) {
	type args struct {
		part      Part
		workflows map[string][]WorkflowStep
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPartAccepted(tt.args.part, tt.args.workflows); got != tt.want {
				t.Errorf("isPartAccepted() = %v, want %v", got, tt.want)
			}
		})
	}
}
