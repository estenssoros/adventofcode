package day3

import (
	"fmt"
	"testing"
)

func TestCalculateTrees(t *testing.T) {
	input, err := readInput()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(calculateTrees(input, slope{3, 1}))
}

func TestMultiplySlopes(t *testing.T) {
	input, err := readInput()
	if err != nil {
		t.Fatal(err)
	}
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	out := 1
	for _, slope := range slopes {
		out *= calculateTrees(input, slope)
	}
	fmt.Println(out)
}
