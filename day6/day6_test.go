package day6

import (
	"fmt"
	"testing"

	"github.com/estenssoros/adventofcode/helpers"
)

func TestPart1(t *testing.T) {
	input, err := helpers.ReadInputChunks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part1: ", Part1(input))
}

func TestPart2(t *testing.T) {
	input, err := helpers.ReadInputChunks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part2: ", Part2(input))
}
