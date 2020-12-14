package day10

import (
	"fmt"
	"testing"

	"github.com/estenssoros/adventofcode/helpers"
)

func TestPart1(t *testing.T) {
	input, err := helpers.ReadInputInt()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part1: ", part1(input))
}

func TestPart2(t *testing.T) {
	input, err := helpers.ReadInputInt()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(part2(input))
}
