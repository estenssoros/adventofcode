package day3

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

func readInput() ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, errors.Wrap(err, "os.Open")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := []string{}
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out, nil
}
