package day8

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/estenssoros/adventofcode/helpers"
	"github.com/pkg/errors"
)

type Step struct {
	Action string
	Val    int
}

func (s Step) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

func ReadInput() ([]*Step, error) {
	actions := []*Step{}
	for input := range helpers.ReadInputChan() {
		if input.Error != nil {
			return nil, input.Error
		}
		action, err := parseStep(input.Val)
		if err != nil {
			return nil, errors.Wrap(err, "parseStep")
		}
		actions = append(actions, action)
	}
	return actions, nil
}

func parseStep(input string) (*Step, error) {
	fields := strings.Fields(input)
	if len(fields) != 2 {
		return nil, errors.Errorf("could not parse: %s", input)
	}
	i, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, errors.Wrap(err, "strconv.Atoi")
	}
	return &Step{
		Action: fields[0],
		Val:    i,
	}, nil
}

func part1(input []*Step) (val int, terminated bool) {
	var idx int
	visited := map[int]bool{}
	for idx < len(input) {
		if _, ok := visited[idx]; ok {
			return val, false
		}
		visited[idx] = true
		step := input[idx]
		switch step.Action {
		case "acc":
			val += step.Val
			idx++
		case "jmp":
			idx += step.Val
		case "nop":
			idx++
		}
	}
	return val, true
}

func part2(steps []*Step) int {
	nopIdx := []int{}
	jumpIdx := []int{}
	for idx, step := range steps {
		switch step.Action {
		case "nop":
			nopIdx = append(nopIdx, idx)
		case "jmp":
			jumpIdx = append(jumpIdx, idx)
		}
	}
	for _, idx := range nopIdx {
		steps[idx].Action = "jmp"
		if val, terminated := part1(steps); terminated {
			return val
		}
		steps[idx].Action = "nop"
	}
	for _, idx := range jumpIdx {
		steps[idx].Action = "nop"
		if val, terminated := part1(steps); terminated {
			return val
		}
		steps[idx].Action = "jmp"
	}
	return -1
}
