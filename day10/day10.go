package day10

import (
	"encoding/json"
	"sort"
)

func part1(input []int) int {
	sort.Ints(input)
	var (
		jolts  int
		count1 int
		count3 int
	)
	for _, adapater := range input {
		delta := adapater - jolts
		if delta < 1 {
			continue
		} else if delta > 3 {
			return count1 * count3
		}
		jolts = adapater
		switch delta {
		case 1:
			count1++
		case 3:
			count3++
		default:
		}
	}
	count3++
	return count1 * count3
}

func part2(input []int) int {
	sort.Ints(input)
	return part2Helper(option{0, 0}, 0, input, map[int]int{})
}

func part2Helper(opt option, jolts int, input []int, cache map[int]int) int {
	if count, ok := cache[opt.Adapter]; ok {
		return count
	}
	options := getValidAdapters(opt.Idx, jolts, input)
	if len(options) == 0 {
		return 1
	}
	var count int
	for _, opt = range options {
		cache[opt.Adapter] = part2Helper(opt, opt.Adapter, input, cache)
		count += cache[opt.Adapter]
	}
	return count
}

type option struct {
	Adapter int
	Idx     int
}

func (o option) String() string {
	ju, _ := json.Marshal(o)
	return string(ju)
}

func getValidAdapters(pos, jolts int, input []int) []option {
	options := []option{}
	for i := pos; i < len(input); i++ {
		test := input[i] - jolts
		if test > 0 && test < 4 {
			options = append(options, option{input[i], i})
		}
		if test > 3 {
			return options
		}
	}
	return options
}
