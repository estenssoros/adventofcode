package day9

import "math"

func part1(input []int, preamble int) int {
	nums := make([]int, preamble)
	for i := 0; i < preamble; i++ {
		nums[i] = input[i]
	}
	for i := preamble; i < len(input); i++ {
		if !isSum(input[i], nums) {
			return input[i]
		}
		nums = append(nums, input[i])
		nums = nums[1:]
	}
	return -1
}

func isSum(target int, nums []int) bool {
	complements := map[int]int{}
	for _, el := range nums {
		_, ok := complements[target-el]
		if ok {
			return true
		}
		complements[el] = el
	}
	return false
}

func part2(input []int, preamble int) int {
	target := part1(input, preamble)
	nums := findContiguousSum(target, input)
	return min(nums) + max(nums)
}

func findContiguousSum(target int, input []int) []int {
	for i := 0; i < len(input); i++ {
		nums, ok := findContiguousSumFromIdx(i, target, input)
		if ok {
			return nums
		}
	}
	return []int{}
}

func findContiguousSumFromIdx(idx, targetSum int, input []int) ([]int, bool) {
	var sum int
	vals := []int{}
	for i := idx; i < len(input); i++ {
		sum += input[i]
		if sum <= targetSum {
			vals = append(vals, input[i])
		} else if sum > targetSum {
			return vals, false
		}
		if sum == targetSum {
			return vals, true
		}
	}
	return vals, false
}

func min(nums []int) int {
	min := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
func max(nums []int) int {
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
