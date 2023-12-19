package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	timeStart := time.Now()

	fmt.Println("--- Day 19: Aplenty ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

const (
	LessThan    = "<"
	GreaterThan = ">"
	Accepted    = "A"
	Rejected    = "R"
)

type WorkflowStep struct {
	Field          string
	Condition      string
	Value          int
	NextWorkflowId string
}

type Part struct {
	X int
	M int
	A int
	S int
}

func partOne(input string) int {
	inputs := strings.Split(input, "\n\n")
	inputWorkflows := strings.Split(inputs[0], "\n")
	inputParts := strings.Split(inputs[1], "\n")

	workflowIdRegex := regexp.MustCompile(`(?P<WorkflowId>\w+){(?P<WorkflowSteps>.+)}`)
	workflowStepRegex := regexp.MustCompile(`(?P<Field>[xmas])(?P<Condition>[<>])(?P<Value>\d+):(?P<Result>\w+)`)

	workflows := make(map[string][]WorkflowStep, len(inputWorkflows))

	for _, workflowLine := range inputWorkflows {
		workflowIdMatches := workflowIdRegex.FindStringSubmatch(workflowLine)
		workflowId := workflowIdMatches[workflowIdRegex.SubexpIndex("WorkflowId")]
		workflowSteps := strings.Split(workflowIdMatches[workflowIdRegex.SubexpIndex("WorkflowSteps")], ",")

		flows := make([]WorkflowStep, len(workflowSteps))

		for i, step := range workflowSteps {
			if strings.Contains(step, ":") {
				stepMatches := workflowStepRegex.FindStringSubmatch(step)
				value, _ := strconv.Atoi(stepMatches[workflowStepRegex.SubexpIndex("Value")])

				flows[i] = WorkflowStep{
					Field:          stepMatches[workflowStepRegex.SubexpIndex("Field")],
					Condition:      stepMatches[workflowStepRegex.SubexpIndex("Condition")],
					Value:          value,
					NextWorkflowId: stepMatches[workflowStepRegex.SubexpIndex("Result")],
				}
			} else {
				flows[i].NextWorkflowId = step
			}
		}

		workflows[workflowId] = flows
	}

	sumOfAcceptedParts := 0

	for _, inputPart := range inputParts {
		partNumbers := util.ReadNumbersInString(inputPart)

		if len(partNumbers) != 4 {
			panic("invalid part" + inputPart)
		}

		part := Part{
			X: partNumbers[0],
			M: partNumbers[1],
			A: partNumbers[2],
			S: partNumbers[3],
		}

		if isPartAccepted(part, workflows) {
			sumOfAcceptedParts += util.SumOfSlice(partNumbers)
		}
	}

	return sumOfAcceptedParts
}

func (w WorkflowStep) isNextStepFinal() bool {
	return w.NextWorkflowId == Accepted || w.NextWorkflowId == Rejected
}

func isPartAccepted(part Part, workflows map[string][]WorkflowStep) bool {
	workflowStepId := 0
	currentWorkflowId := "in"

	for {
		currentWorkflowStep := workflows[currentWorkflowId][workflowStepId]

		if part.matchesCondition(currentWorkflowStep.Field, currentWorkflowStep.Condition, currentWorkflowStep.Value) {
			workflowStepId = 0
			currentWorkflowId = currentWorkflowStep.NextWorkflowId

			if currentWorkflowStep.isNextStepFinal() {
				break
			}
		} else {
			workflowStepId++
		}
	}

	fmt.Println("returning final step", currentWorkflowId)

	return currentWorkflowId == Accepted
}

func (p Part) matchesCondition(field string, condition string, value int) bool {
	if condition != LessThan && condition != GreaterThan {
		return true
	}

	inputValue := 0

	if field == "x" {
		inputValue = p.X
	} else if field == "m" {
		inputValue = p.M
	} else if field == "a" {
		inputValue = p.A
	} else if field == "s" {
		inputValue = p.S
	}

	if condition == LessThan {
		return inputValue < value
	}

	if condition == GreaterThan {
		return inputValue > value
	}

	panic("help")
}

func partTwo(input string) int {
	return 0
}
