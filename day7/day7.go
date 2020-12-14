package day7

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/estenssoros/adventofcode/helpers"
	"github.com/pkg/errors"
)

type Bag struct {
	Amount   int
	Pattern  string
	Color    string
	Children []*Bag
}

func (b *Bag) Equals(other *Bag) bool {
	return b.Pattern == other.Pattern && b.Color == other.Color
}

func (b Bag) String() string {
	ju, _ := json.MarshalIndent(b, "", " ")
	return string(ju)
}

func (b Bag) Key() BagKey {
	return BagKey{
		Pattern: b.Pattern,
		Color:   b.Color,
	}
}

type BagKey struct {
	Pattern string
	Color   string
}

func (k BagKey) String() string {
	return fmt.Sprintf("%s-%s", k.Pattern, k.Color)
}

var parentBag = regexp.MustCompile(`(\w+) (\w+) bags contain`)
var childrenBags = regexp.MustCompile(`(\d+) (\w+) (\w+) bag`)

func ReadInput() ([]*Bag, error) {
	bags := []*Bag{}
	for input := range helpers.ReadInputChan() {
		if input.Error != nil {
			return nil, input.Error
		}
		bag, err := matchBag(input.Val)
		if err != nil {
			return nil, errors.Wrap(err, "matchbag")
		}
		bags = append(bags, bag)
	}
	return bags, nil
}

func matchBag(input string) (*Bag, error) {
	parentBagMatch := parentBag.FindStringSubmatch(input)
	if len(parentBagMatch) != 3 {
		return nil, errors.Errorf("could not match: %s", input)
	}
	bag := &Bag{
		Amount:  1,
		Pattern: parentBagMatch[1],
		Color:   parentBagMatch[2],
	}
	childrenBagMatch := childrenBags.FindAllStringSubmatch(input, -1)
	for _, childBagMatch := range childrenBagMatch {
		childBag, err := matchChildBag(childBagMatch)
		if err != nil {
			return nil, errors.Wrap(err, "matchChildBag")
		}
		bag.Children = append(bag.Children, childBag)
	}
	return bag, nil
}

func matchChildBag(match []string) (*Bag, error) {
	if len(match) != 4 {
		return nil, errors.Errorf("could not match: %s", match[0])
	}
	i, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, errors.Wrap(err, "strconv.Atoi")
	}
	return &Bag{
		Amount:  i,
		Pattern: match[2],
		Color:   match[3],
	}, nil
}

func createBagTree(bags []*Bag) map[BagKey]*Bag {
	bagTree := map[BagKey]*Bag{}
	for _, bag := range bags {
		bagTree[bag.Key()] = bag
	}
	return bagTree
}

func part1(bags []*Bag) int {
	bag := &Bag{
		Pattern: "shiny",
		Color:   "gold",
	}
	return part1Helper(bag, createBagTree(bags), map[BagKey]int{})
}

func part1Helper(myBag *Bag, bagTree map[BagKey]*Bag, cache map[BagKey]int) int {
	var count int
	for bagKey, bag := range bagTree {
		_, visited := cache[bagKey]
		if isInChildren(myBag, bag.Children) && !visited {
			cache[bagKey] = 1
			count++
			count += part1Helper(bagTree[bagKey], bagTree, cache)
		}
	}
	return count
}

func isInChildren(myBag *Bag, bags []*Bag) bool {
	for _, bag := range bags {
		if bag.Equals(myBag) {
			return true
		}
	}
	return false
}

type visitor struct {
	BagKey     BagKey
	Multiplier int
}

func part2(bags []*Bag) int {
	bag := &Bag{
		Pattern: "shiny",
		Color:   "gold",
	}
	bagTree := createBagTree(bags)
	node, ok := bagTree[bag.Key()]
	if !ok {
		panic("could not locate bag")
	}
	if len(node.Children) == 0 {
		return 1
	}
	var count int
	toVisit := []visitor{}
	for _, child := range node.Children {
		toVisit = append(toVisit, visitor{
			BagKey:     child.Key(),
			Multiplier: child.Amount,
		})
	}
	for len(toVisit) > 0 {
		visiting := toVisit[0]
		toVisit = toVisit[1:]
		bag, ok := bagTree[visiting.BagKey]
		if !ok {
			panic("could not locate bag")
		}
		count += bag.Amount * visiting.Multiplier
		for _, child := range bag.Children {
			toVisit = append(toVisit, visitor{
				BagKey:     child.Key(),
				Multiplier: visiting.Multiplier * child.Amount,
			})
		}
	}
	return count
}
