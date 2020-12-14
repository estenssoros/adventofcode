package day9

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
	fmt.Println("part 1: ", part1(input, 25))
}
func TestPart2(t *testing.T) {
	input, err := helpers.ReadInputInt()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part 2: ", part2(input, 25))
}
