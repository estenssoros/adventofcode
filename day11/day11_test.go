package day11

import (
	"fmt"
	"strings"
	"testing"

	"github.com/estenssoros/adventofcode/helpers"
	"gopkg.in/go-playground/assert.v1"
)

func TestPart1(t *testing.T) {
	input, err := helpers.ReadInputRune()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part1: ", part1(input))
}

func TestPart2(t *testing.T) {
	input, err := helpers.ReadInputRune()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("part2: ", part2(input))
}

var testInput = []string{`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,
	`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,
	`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`,
	`#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`,
	`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`,
	`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`,
	`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`,
}

// func TestPart2Input(t *testing.T) {
// 	// assert.Equal(t, 26, part2(parseRaw(testInput)))
// 	next := nextStatePart2(parseRaw(testInput[0]))
// 	for i := 1; i < len(testInput); i++ {
// 		fmt.Println(i)
// 		if want, have := countState(parseRaw(testInput[i]), occupied), countState(next, occupied); want != have {
// 			fmt.Println("want:")
// 			printState(parseRaw(testInput[i]))
// 			fmt.Println("")
// 			fmt.Println("have:")
// 			printState(next)
// 			t.Fatal("asdf")
// 		}
// 		printState(parseRaw(testInput[i]))
// 		fmt.Println("")
// 		next = nextStatePart2(next)
// 	}
// }

func parseRaw(input string) [][]rune {
	output := [][]rune{}
	for _, row := range strings.Split(input, "\n") {
		output = append(output, []rune(row))
	}
	return output
}

func TestCount(t *testing.T) {
	raw := `.......#.
	...#.....
	.#.......
	.........
	..#L....#
	....#....
	.........
	#........
	...#.....`
	assert.Equal(t, 8, countDirectionOccupied(4, 3, parseRaw(raw)))
	raw = `.............
	.L.L.#.#.#.#.
	.............`
	assert.Equal(t, 0, countDirectionOccupied(1, 1, parseRaw(raw)))
	raw = `.##.##.
	#.#.#.#
	##...##
	...L...
	##...##
	#.#.#.#
	.##.##.`
	assert.Equal(t, 0, countDirectionOccupied(3, 3, parseRaw(raw)))
	raw = `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`
	assert.Equal(t, 1, countDirectionOccupied(0, 2, parseRaw(raw)))
	assert.Equal(t, false, noDirectionOccupied(0, 2, parseRaw(raw)))
}
