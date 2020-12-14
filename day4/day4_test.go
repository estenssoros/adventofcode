package day4

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestValdiatePassport(t *testing.T) {
	input, err := getInput()
	if err != nil {
		t.Fatal(err)
	}
	var count int
	for _, l := range input {
		if isValidPassport(l, requiredFields) {
			count++
		}
	}
	fmt.Println(count)
}

func TestFieldIsValid(t *testing.T) {
	//eyr:2025 hgt:161cm iyr:1962 pid:394421140 ecl:gry cid:209 hcl:#efcc98 byr:2001
	assert.Equal(t, true, field("eyr").isValid("2025"))
	assert.Equal(t, true, field("hgt").isValid("161cm"))
	assert.Equal(t, false, field("iyr").isValid("1962"))
	assert.Equal(t, true, field("pid").isValid("394421140"))
	assert.Equal(t, true, field("ecl").isValid("gry"))
	assert.Equal(t, true, field("cid").isValid("209"))
	assert.Equal(t, true, field("hcl").isValid("#efcc98"))
	assert.Equal(t, true, field("byr").isValid("2001"))
	assert.Equal(t, true, 'e' >= 'a' && 'e' <= 'f')
}
