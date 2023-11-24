package worm

import (
	"github.com/creepitall/goworm/internal/models"
)

type Area interface {
	IsInside(p models.Position) bool
}

type Worm struct {
	positions models.Positions
	area      Area
}

func New(area Area) *Worm {
	return &Worm{
		area:      area,
		positions: models.Positions{{X: 1, Y: 1}},
	}
}

func (w *Worm) Change(way models.Way) bool {
	w.change(way)

	return w.area.IsInside(w.GetHead())
}

func (w *Worm) change(way models.Way) {
	head := w.positions[len(w.positions)-1].Change(way)
	tail := w.positions[1:]

	out := make(models.Positions, 0, len(w.positions))
	out = append(out, tail...)
	out = append(out, head)

	w.positions = out
}

func (w *Worm) Add(way models.Way) {
	w.positions = append(w.positions, w.positions[len(w.positions)-1].Change(way))
}

func (w Worm) GetHead() models.Position {
	return w.positions[len(w.positions)-1]
}

func (w Worm) Positions() models.Positions {
	return w.positions
}
