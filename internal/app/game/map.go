package game

import "github.com/creepitall/goworm/internal/models"

type Map struct {
	width  int
	height int
	X      int
	Y      int
}

func (m *Map) Init(w, h int) *Map {
	m.width = w
	m.height = h
	m.X = m.width / 40
	m.Y = m.height / 40
	return m
}

func (m Map) IsOutside(p models.Position) bool {
	x, y := p.Get()
	if x > m.X || y > m.Y {
		return true
	}
	if x < 0 || y < 0 {
		return true
	}
	return false
}
