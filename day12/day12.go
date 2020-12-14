package day12

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

const (
	North = 0
	East  = 90
	South = 180
	West  = 270
)

type Ferry struct {
	Direction int
	Point
	WayPoint Point
}

type Point struct {
	X int
	Y int
}

func (p *Point) Rotate(origin Point, mag int, rotateFunc func(float64) float64) {
	rad := rotateFunc(float64(mag) * math.Pi / 180)
	temp := p.X
	p.X = int(math.Round(float64(p.X)*math.Cos(rad) - float64(p.Y)*math.Sin(rad)))
	p.Y = int(math.Round(float64(temp)*math.Sin(rad) + float64(p.Y)*math.Cos(rad)))
}

func (p *Point) RotateLeft(origin Point, mag int) {
	p.Rotate(origin, mag, func(i float64) float64 { return i })
}

func (p *Point) RotateRight(origin Point, mag int) {
	p.Rotate(origin, mag, func(i float64) float64 { return -1 * i })
}

func (f Ferry) String() string {
	ju, _ := json.MarshalIndent(f, "", " ")
	return string(ju)
}

func (f *Ferry) Process1(input string) error {
	action := input[0]
	mag, err := strconv.Atoi(input[1:])
	if err != nil {
		return errors.Wrap(err, "strconv.Atoi")
	}
	switch action {
	case 'N':
		f.Y += mag
	case 'S':
		f.Y -= mag
	case 'E':
		f.X += mag
	case 'W':
		f.X -= mag
	case 'L':
		f.Direction = turnLeft(f.Direction, mag)
	case 'R':
		f.Direction = turnRight(f.Direction, mag)
	case 'F':
		f.moveForward(f.Direction, mag)
	default:
		return errors.Errorf("unknown action: %s", string(action))
	}
	return nil
}

func turnLeft(direction int, mag int) int {
	newDirection := direction - mag
	if newDirection < 0 {
		newDirection = 360 + newDirection
	}
	return newDirection
}

func turnRight(direction, mag int) int {
	newDirection := direction + mag
	if newDirection >= 360 {
		newDirection = newDirection - 360
	}
	return newDirection
}

func (p *Point) moveForward(direction, mag int) {
	switch direction {
	case North:
		p.Y += mag
	case South:
		p.Y -= mag
	case East:
		p.X += mag
	case West:
		p.X -= mag
	}
}

func part1(input []string) (int, error) {
	ferry := Ferry{Direction: East}
	for _, direction := range input {
		if err := ferry.Process1(direction); err != nil {
			return -1, errors.Wrap(err, "ferry.Process")
		}
	}
	return manhattanDiff(ferry.X, ferry.Y), nil

}

func manhattanDiff(x1, y1 int) int {
	return abs(x1) + abs(y1)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func part2(input []string) (int, error) {
	ferry := Ferry{
		Direction: East,
		WayPoint:  Point{10, 1},
	}
	for _, i := range input {
		if err := ferry.Process2(i); err != nil {
			return 0, errors.Wrap(err, "")
		}
	}
	fmt.Println(ferry)
	return manhattanDiff(ferry.X, ferry.Y), nil
}

func (f *Ferry) Process2(input string) error {
	action := input[0]
	mag, err := strconv.Atoi(input[1:])
	if err != nil {
		return errors.Wrap(err, "strconv.Atoi")
	}
	switch action {
	case 'N':
		f.WayPoint.Y += mag
	case 'S':
		f.WayPoint.Y -= mag
	case 'E':
		f.WayPoint.X += mag
	case 'W':
		f.WayPoint.X -= mag
	case 'L':
		f.WayPoint.RotateLeft(f.Point, mag)
	case 'R':
		f.WayPoint.RotateRight(f.Point, mag)
	case 'F':
		f.X += mag * f.WayPoint.X
		f.Y += mag * f.WayPoint.Y
	default:
		return errors.Errorf("unknown action: %s", string(action))
	}
	return nil
}
