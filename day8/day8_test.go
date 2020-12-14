package day8

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	steps, err := ReadInput()
	if err != nil {
		t.Fatal(err)
	}
	val, _ := part1(steps)
	fmt.Println("part: ", val)
}
func TestPart2(t *testing.T) {
	steps, err := ReadInput()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part2: ", part2(steps))
}
