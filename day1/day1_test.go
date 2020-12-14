package day1

import (
	"fmt"
	"testing"

	"github.com/estenssoros/adventofcode/helpers"
)

func TestGetInput(t *testing.T) {
	out, err := helpers.ReadInputInt()
	if err != nil {
		t.Fatal(err)
	}
	twoNum, err := twoNumberSum(out, 2020)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("two: ", twoNum.a*twoNum.b)
}

func TestThreeNumberSum(t *testing.T) {
	input, err := helpers.ReadInputInt()
	if err != nil {
		t.Fatal(err)
	}
	threeNum, err := threeNumberSum(input, 2020)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("three: ", threeNum.a*threeNum.b*threeNum.c)
}
