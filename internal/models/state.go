package models

import (
	"strconv"
	"strings"
)

type Position struct {
	X  int
	Y  int
	XY string
}

func (p Position) Get() (int, int) {
	return p.X, p.Y
}

func (p Position) Change(w Way) Position {
	x, y := p.Get()

	switch w {
	case Right:
		x = p.X + 1
	case Left:
		x = p.X - 1
	case Up:
		y = p.Y - 1
	case Down:
		y = p.Y + 1
	}
	return Position{
		X:  x,
		Y:  y,
		XY: strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, "_"),
	}
}

type Positions []Position

type Way string

const (
	Right Way = "right"
	Left  Way = "left"
	Up    Way = "up"
	Down  Way = "down"
)

func (w Way) Reverse() Way {
	switch w {
	case Right:
		return Left
	case Left:
		return Right
	case Up:
		return Down
	case Down:
		return Up
	default:
		return Up
	}
}

func GetWayFromString(s string) Way {
	switch s {
	case "right":
		return Right
	case "left":
		return Left
	case "up":
		return Up
	case "down":
		return Down
	default:
		return Up
	}
}
