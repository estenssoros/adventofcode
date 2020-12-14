package day5

import (
	"encoding/json"
	"sort"

	"github.com/estenssoros/adventofcode/helpers"
	"github.com/pkg/errors"
)

var maxRow = 127
var maxSeats = 7

type Seat struct {
	Row int
	Col int
}

func (s Seat) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

func (s Seat) ID() int {
	return s.Row*8 + s.Col
}

type Seats []*Seat

func (s Seats) Sort() {
	sort.Slice(s, func(i, j int) bool {
		if s[i].Row != s[j].Row {
			return s[i].Row < s[j].Row
		}
		return s[i].Col < s[j].Col
	})
}

func ParseSeat(input string) *Seat {
	return &Seat{
		Row: parseRow(input[:7]),
		Col: parseCol(input[7:]),
	}
}

func parseRow(input string) int {
	left, right := 0, maxRow
	for _, r := range input {
		if r == 'F' {
			right = right - (right-left)/2 - 1
		} else if r == 'B' {
			left = left + (right-left)/2 + 1
		}
	}
	return left
}

func parseCol(input string) int {
	left, right := 0, maxSeats
	for _, r := range input {
		if r == 'L' {
			right = right - (right-left)/2 - 1
		} else if r == 'R' {
			left = left + (right-left)/2 + 1
		}
	}
	return left
}

func getSeats() ([]*Seat, error) {
	input, err := helpers.ReadInput()
	if err != nil {
		return nil, errors.Wrap(err, "helpers.ReadInput")
	}
	seats := []*Seat{}
	for _, row := range input {
		seats = append(seats, ParseSeat(row))
	}
	return seats, nil
}

func getMaxSeatID(seats []*Seat) int {
	var max int
	for _, seat := range seats {
		if seatID := seat.ID(); seatID > max {
			max = seatID
		}
	}
	return max
}

func generateAllSeats() []*Seat {
	seats := []*Seat{}
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxSeats; col++ {
			seats = append(seats, &Seat{Row: row, Col: col})
		}
	}
	return seats
}
