package day6

import "fmt"

func Part1(input [][]string) int {
	var count int
	for _, group := range input {
		groupCounter := map[rune]struct{}{}
		for _, person := range group {
			for _, r := range person {
				groupCounter[r] = struct{}{}
			}
		}
		count += len(groupCounter)
	}
	return count
}

func Part2(input [][]string) int {
	var count int
	for _, group := range input {
		groupCounter := map[rune]int{}
		for _, person := range group {
			for _, r := range person {
				groupCounter[r]++
			}
		}
		fmt.Println(group)
		fmt.Println(groupCounter)
		fmt.Println(len(group), countHas(groupCounter, len(group)))
		count += countHas(groupCounter, len(group))
	}
	return count
}

func countHas(counter map[rune]int, target int) int {
	var count int
	for _, test := range counter {
		if test == target {
			count++
		}
	}
	return count
}
