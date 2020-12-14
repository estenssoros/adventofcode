package day7

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestMatchBag(t *testing.T) {
	bag, err := matchBag("plaid bronze bags contain 5 mirrored orange bags, 4 plaid cyan bags, 1 dotted black bag.")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, bag.Amount)
	assert.Equal(t, "plaid", bag.Pattern)
	assert.Equal(t, "bronze", bag.Color)
	assert.Equal(t, 3, len(bag.Children))
}

func TestPart1(t *testing.T) {
	bags, err := ReadInput()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part1: ", part1(bags))
}

func TestPart2(t *testing.T) {
	bags, err := ReadInput()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part2: ", part2(bags))
}
