package day4

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func getInput() ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, errors.Wrap(err, "os.Open")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := []string{}
	var batch []string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			out = append(out, strings.Join(batch, " "))
			batch = nil
			continue
		}
		batch = append(batch, text)
	}
	out = append(out, strings.Join(batch, " "))
	return out, nil
}
