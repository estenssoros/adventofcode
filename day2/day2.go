package day2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

type password struct {
	min   int
	max   int
	char  rune
	input string
}

func (p *password) isValid1() bool {
	var count int
	for _, r := range p.input {
		if r == p.char {
			count++
		}
		if count > p.max {
			return false
		}
	}
	return count >= p.min
}

func (p *password) isValid2() bool {
	char1 := rune(p.input[p.min-1])
	char2 := rune(p.input[p.max-1])
	if char1 != p.char && char2 != p.char {
		return false
	}
	return !(char1 == p.char && char2 == p.char)
}

var myExp = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>\w): (?P<input>\w+)`)

func lineToPassword(input string) (*password, error) {
	match := myExp.FindStringSubmatch(input)
	if len(match) != 5 {
		return nil, errors.Errorf("could not match: [%s]: %v", input, match)
	}
	p := &password{}
	{
		i, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, errors.Wrap(err, "strconv.Atoi")
		}
		p.min = i
	}
	{
		i, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, errors.Wrap(err, "strconv.Atoi")
		}
		p.max = i
	}
	p.char = rune(match[3][0])
	p.input = match[4]
	return p, nil
}

func getInput() ([]*password, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := []*password{}
	for scanner.Scan() {
		p, err := lineToPassword(scanner.Text())
		if err != nil {
			return nil, errors.Wrap(err, "lineToPassword")
		}
		out = append(out, p)
	}
	return out, nil
}

func countValidPasswords1(passwords []*password) (int, error) {
	var count int
	for _, password := range passwords {
		if password.isValid1() {
			count++
		}
	}
	return count, nil
}

func countValidPasswords2(passwords []*password) (int, error) {
	var count int
	for _, password := range passwords {
		if password.isValid2() {
			count++
		}
	}
	return count, nil
}
