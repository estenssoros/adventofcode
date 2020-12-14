package helpers

import (
	"bufio"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type StringInput struct {
	Val   string
	Error error
}

func ReadInputChan() chan *StringInput {
	ch := make(chan *StringInput)
	go func() {
		defer close(ch)
		f, err := os.Open("input.txt")
		if err != nil {
			ch <- &StringInput{Error: err}
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- &StringInput{Val: scanner.Text()}
		}
	}()
	return ch
}

func ReadInput() ([]string, error) {
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

func ReadInputChunks() ([][]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, errors.Wrap(err, "os.Open")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := [][]string{}
	chunk := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			out = append(out, chunk)
			chunk = nil
			continue
		}
		chunk = append(chunk, text)
	}
	out = append(out, chunk)
	return out, nil

}

func ReadInputInt() ([]int, error) {
	out := []int{}
	for input := range ReadInputChan() {
		if input.Error != nil {
			return nil, input.Error
		}
		i, err := strconv.Atoi(input.Val)
		if err != nil {
			return nil, errors.Wrap(err, "strconv.Atoi")
		}
		out = append(out, i)
	}
	return out, nil
}

func ReadInputRune() ([][]rune, error) {
	out := [][]rune{}
	for input := range ReadInputChan() {
		if input.Error != nil {
			return nil, input.Error
		}
		out = append(out, []rune(input.Val))
	}
	return out, nil
}
