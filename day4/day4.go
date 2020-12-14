package day4

import (
	"strconv"
	"strings"
)

type field string

var (
	BirthYear      field = "byr"
	IssueYear      field = "iyr"
	ExpirationYear field = "eyr"
	Height         field = "hgt"
	HairColor      field = "hcl"
	EyeColor       field = "ecl"
	PassportID     field = "pid"
	CountryID      field = "cid"
)

func (f field) isValid(input string) bool {
	switch f {
	case BirthYear:
		return isDigitLength(input, 4) && isBetween(1920, 2002, input)
	case IssueYear:
		return isDigitLength(input, 4) && isBetween(2010, 2020, input)
	case ExpirationYear:
		return isDigitLength(input, 4) && isBetween(2020, 2030, input)
	case Height:
		return isValidHeight(input)
	case HairColor:
		return isValidHair(input)
	case EyeColor:
		return isValidEye(input)
	case PassportID:
		return isDigitLength(input, 9)
	case CountryID:
		return true
	}
	return false
}

func isValidHair(input string) bool {
	if input[0] != '#' {
		return false
	}
	if len(input) != 7 {
		return false
	}
	for i := 1; i < len(input); i++ {
		val := rune(input[i])
		if !isCharBetween('0', '9', val) && !isCharBetween('a', 'f', val) {
			return false
		}
	}
	return true
}
func isValidEye(input string) bool {
	switch input {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func isCharBetween(min, max, val rune) bool {
	return val >= min && val <= max
}

func isDigitLength(input string, length int) bool {
	if len(input) != length {
		return false
	}
	_, err := strconv.Atoi(input)
	return err == nil
}

func isBetween(min, max int, input string) bool {
	i, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}

func isValidHeight(input string) bool {
	if strings.HasSuffix(input, "cm") {
		return isBetween(150, 193, input[:len(input)-2])
	}
	if strings.HasSuffix(input, "in") {
		return isBetween(59, 76, input[:len(input)-2])
	}
	return false

}

var requiredFields = []field{
	BirthYear,
	IssueYear,
	ExpirationYear,
	Height,
	HairColor,
	EyeColor,
	PassportID,
}

func isValidPassport(input string, requiredFields []field) bool {
	fields := strings.Fields(input)
	if len(fields) < len(requiredFields) {
		return false
	}
	has := map[string]string{}
	for _, field := range fields {
		splits := strings.Split(field, ":")
		if len(splits) != 2 {
			return false
		}
		has[splits[0]] = splits[1]
	}
	for _, requiredField := range requiredFields {
		input, ok := has[string(requiredField)]
		if !ok {
			return false
		}
		if !requiredField.isValid(input) {
			return false
		}
	}
	return true
}
