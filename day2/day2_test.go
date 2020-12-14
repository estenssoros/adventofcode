package day2

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCountValidPasswords1(t *testing.T) {
	passwords, err := getInput()
	if err != nil {
		t.Fatal(err)
	}
	count, err := countValidPasswords1(passwords)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("1:", count)
}
func TestCountValidPasswords2(t *testing.T) {
	passwords, err := getInput()
	if err != nil {
		t.Fatal(err)
	}
	count, err := countValidPasswords2(passwords)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("2:", count)
}

func TestAPassword(t *testing.T) {
	p := &password{
		min:   1,
		max:   3,
		char:  'a',
		input: "abcde",
	}
	assert.Equal(t, true, p.isValid1())
}

func TestValid2(t *testing.T) {
	{
		p := &password{
			min:   1,
			max:   3,
			char:  'a',
			input: "abcde",
		}
		assert.Equal(t, true, p.isValid2())
	}
	{
		p := &password{
			min:   1,
			max:   3,
			char:  'b',
			input: "cdefg",
		}
		assert.Equal(t, false, p.isValid2())
	}
	{
		p := &password{
			min:   2,
			max:   9,
			char:  'c',
			input: "ccccccccc",
		}
		assert.Equal(t, false, p.isValid2())
	}
}
