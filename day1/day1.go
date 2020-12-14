package day1

import (
	"sort"

	"github.com/pkg/errors"
)

type twoSums struct {
	a int
	b int
}

func twoNumberSum(numbers []int, target int) (*twoSums, error) {
	complements := map[int]int{}
	for _, el := range numbers {
		a, ok := complements[target-el]
		if ok {
			return &twoSums{a, el}, nil
		}
		complements[el] = el
	}
	return nil, errors.New("could not find target")
}

type threeSums struct {
	a, b, c int
}

func threeNumberSum(numbers []int, target int) (*threeSums, error) {
	sort.Ints(numbers)
	for i := 0; i < len(numbers)-2; i++ {
		left, right := i+1, len(numbers)-1
		for left < right {
			currentSum := numbers[i] + numbers[left] + numbers[right]
			if currentSum == target {
				return &threeSums{numbers[i], numbers[left], numbers[right]}, nil
			} else if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return nil, errors.New("could not find target")
}
