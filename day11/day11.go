package day11

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
)

func part1(input [][]rune) int {
	var prev int
	for {
		input = nextStatePart1(input)
		if next := countState(input, occupied); next == prev {
			return next
		} else {
			prev = next
		}
	}
}

func nextStatePart1(input [][]rune) [][]rune {
	out := copyState(input)
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			switch input[row][col] {
			case empty:
				if noAdjacentOccupied(row, col, input) {
					out[row][col] = occupied
				}
			case occupied:
				if countAdjacentOccupied(row, col, input) >= 4 {
					out[row][col] = empty
				}
			default:
				out[row][col] = input[row][col]
			}
		}
	}
	return out
}

func noAdjacentOccupied(row, col int, input [][]rune) bool {
	return countAdjacentOccupied(row, col, input) == 0
}

func countAdjacentOccupied(row, col int, input [][]rune) int {
	/*
		x x x
		x * x
		x x x
	*/
	var count int
	if col > 0 {
		if input[row][col-1] == occupied {
			count++
		}
	}
	if row > 0 {
		if input[row-1][col] == occupied {
			count++
		}
		if col > 0 {
			if input[row-1][col-1] == occupied {
				count++
			}
		}
		if col < len(input[row])-1 {
			if input[row-1][col+1] == occupied {
				count++
			}
		}
	}
	if row < len(input)-1 {
		if input[row+1][col] == occupied {
			count++
		}
		if col > 0 {
			if input[row+1][col-1] == occupied {
				count++
			}
		}
		if col < len(input[row])-1 {
			if input[row+1][col+1] == occupied {
				count++
			}
		}
	}

	if col < len(input[row])-1 {
		if input[row][col+1] == occupied {
			count++
		}
	}

	return count
}

func countState(input [][]rune, state rune) int {
	var count int
	for _, row := range input {
		for _, spot := range row {
			if spot == state {
				count++
			}
		}
	}
	return count
}

func copyState(input [][]rune) [][]rune {
	row, col := len(input), len(input[0])
	out := make([][]rune, row)
	for r := 0; r < row; r++ {
		out[r] = make([]rune, col)
		for c := 0; c < col; c++ {
			out[r][c] = input[r][c]
		}
	}
	return out
}

func part2(input [][]rune) int {
	var prev int
	printState(input)
	for {
		input = nextStatePart2(input)
		fmt.Println(strings.Repeat("-", 50))
		printState(input)
		if next := countState(input, occupied); next == prev {
			return next
		} else {
			prev = next
		}
	}
}

func nextStatePart2(input [][]rune) [][]rune {
	out := copyState(input)
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			switch input[row][col] {
			case empty:
				if noDirectionOccupied(row, col, input) {
					out[row][col] = occupied
				}
			case occupied:
				if countDirectionOccupied(row, col, input) >= 5 {
					out[row][col] = empty
				}
			default:
				out[row][col] = input[row][col]
			}
		}
	}
	return out
}

func noDirectionOccupied(row, col int, input [][]rune) bool {
	return countDirectionOccupied(row, col, input) == 0
}

func countDirectionOccupied(row, col int, input [][]rune) int {
	/*
		x x x
		x * x
		x x x
	*/
	var count int
	if row > 0 {
		for r := row - 1; r >= 0; r-- {
			if input[r][col] != floor {
				count += countOccupied(input[r][col])
				// fmt.Println("counting above", count)
				break
			}
		}
		if col > 0 {
			for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
				if input[r][c] != floor {
					count += countOccupied(input[r][c])
					// fmt.Println("counting left above", count)
					break
				}
			}
		}
		if col < len(input[row])-1 {
			for r, c := row-1, col+1; r >= 0 && c < len(input[r]); r, c = r-1, c+1 {
				if input[r][c] != floor {
					count += countOccupied(input[r][c])
					// fmt.Println("counting right above", count)
					break
				}
			}
		}
	}

	if row < len(input)-1 {
		for r := row + 1; r < len(input); r++ {
			if input[r][col] != floor {
				count += countOccupied(input[r][col])
				// fmt.Println("counting below", count)
				break
			}
		}
		if col > 0 {
			for r, c := row+1, col-1; r < len(input) && c >= 0; r, c = r+1, c-1 {
				if input[r][c] != floor {
					count += countOccupied(input[r][c])
					// fmt.Println("counting left below", count)
					break
				}
			}
		}
		if col < len(input[row])-1 {
			for r, c := row+1, col+1; r < len(input) && c < len(input[row]); r, c = r+1, c+1 {
				if input[r][c] != floor {
					count += countOccupied(input[r][c])
					// fmt.Println("counting right below", count)
					break
				}
			}
		}
	}

	if col > 0 {
		for c := col - 1; c >= 0; c-- {
			if input[row][c] != floor {
				count += countOccupied(input[row][c])
				// fmt.Println("counting left", count)
				break
			}
		}
	}

	if col < len(input[row])-1 {
		for c := col + 1; c < len(input[row]); c++ {
			if input[row][c] != floor {
				count += countOccupied(input[row][c])
				// fmt.Println("counting right", count)
				break
			}
		}
	}

	return count
}

func countOccupied(r rune) int {
	if r == occupied {
		return 1
	}
	return 0
}

func printState(state [][]rune) {
	for _, row := range state {
		for _, r := range row {
			switch r {
			case occupied:
				fmt.Print(color.RedString("#"))
			case empty:
				fmt.Print(color.GreenString("L"))
			default:
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
