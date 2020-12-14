package day5

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestParseRow(t *testing.T) {
	assert.Equal(t, 44, parseRow("FBFBBFF"))
}

func TestParseCol(t *testing.T) {
	assert.Equal(t, 5, parseCol("RLR"))
}

func TestHighestSeat(t *testing.T) {
	seats, err := getSeats()
	if err != nil {
		t.Fatal(err)
	}
	maxSeatID := getMaxSeatID(seats)
	fmt.Println(maxSeatID)
}

func TestFindMySeat(t *testing.T) {
	seats, err := getSeats()
	if err != nil {
		t.Fatal(err)
	}
	possibleSeats := generateAllSeats()
	seatLookup := map[int]bool{}
	for _, seat := range possibleSeats {
		seatLookup[seat.ID()] = false
	}
	for _, occupied := range seats {
		seatLookup[occupied.ID()] = true
	}
	empty := []int{}
	for seat, occupied := range seatLookup {
		if !occupied {
			empty = append(empty, seat)
		}
	}
	options := []int{}
	for _, emptySeatID := range empty {
		prevExists := seatLookup[emptySeatID-1]
		nextExists := seatLookup[emptySeatID+1]
		emptyInList := seatLookup[emptySeatID]
		if prevExists && nextExists && !emptyInList {
			options = append(options, emptySeatID)
		}
	}
	fmt.Println(options)
}
