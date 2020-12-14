package day12

import (
	"fmt"
	"testing"

	"github.com/estenssoros/adventofcode/helpers"
	"gopkg.in/go-playground/assert.v1"
)

func TestPart1(t *testing.T) {
	input, err := helpers.ReadInput()
	if err != nil {
		t.Fatal(err)
	}
	p1, err := part1(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part1: ", p1)
}

func TestPart1WithInput(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	f := Ferry{Direction: East}
	// fmt.Println(f.Direction, f.X, f.Y)
	for _, i := range input {
		if err := f.Process1(i); err != nil {
			t.Fatal(err)
		}
		// fmt.Println(i)
		// fmt.Println(f.Direction, f.X, f.Y)
	}
	assert.Equal(t, 25, manhattanDiff(f.X, f.Y))
}
func TestTurning(t *testing.T) {
	input := []string{
		"R90",
		"R90",
		"R90",
	}
	f := Ferry{Direction: East}
	for _, i := range input {
		if err := f.Process1(i); err != nil {
			t.Fatal(err)
		}
	}
	assert.Equal(t, North, f.Direction)
	f.Process1("R90")
	assert.Equal(t, East, f.Direction)
	f.Process1("L180")
	assert.Equal(t, West, f.Direction)
}

func TestPart2(t *testing.T) {
	input, err := helpers.ReadInput()
	if err != nil {
		t.Fatal(err)
	}
	p2, err := part2(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part2: ", p2)
}

var turns = []struct {
	input string
	point Point
}{
	{"N3", Point{10, 4}},
	{"L90", Point{-4, 10}},
	{"L90", Point{-10, -4}},
	{"L90", Point{4, -10}},
	{"L90", Point{10, 4}},
	{"L180", Point{-10, -4}},
	{"L180", Point{10, 4}},
	{"L270", Point{4, -10}},
	{"L90", Point{10, 4}},
	{"R90", Point{4, -10}},
	{"R90", Point{-10, -4}},
	{"R90", Point{-4, 10}},
	{"R90", Point{10, 4}},
	{"R360", Point{10, 4}},
	{"R270", Point{-4, 10}},
}

func TestWayPoint(t *testing.T) {
	f := Ferry{
		WayPoint: Point{10, 1},
	}
	for _, turn := range turns {
		f.Process2(turn.input)
		assert.Equal(t, turn.point, f.WayPoint)
	}
}

func TestPart2WithInput(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	f := Ferry{
		Direction: East,
		WayPoint:  Point{10, 1},
	}
	for _, i := range input {
		if err := f.Process2(i); err != nil {
			t.Fatal(err)
		}
	}
	assert.Equal(t, 286, manhattanDiff(f.X, f.Y))
}
