package area

import "github.com/creepitall/goworm/internal/models"

type Area struct {
	width  int
	height int
	X      int
	Y      int
}

func New(w, h int) Area {
	a := Area{}

	a.width = w
	a.height = h
	a.X = a.width / 40
	a.Y = a.height / 40
	return a
}

func (a Area) IsInside(p models.Position) bool {
	x, y := p.Get()
	if x > a.X || y > a.Y {
		return false
	}
	if x < 0 || y < 0 {
		return false
	}
	return true
}

func (a Area) Get() (int, int) {
	return a.X, a.Y
}
