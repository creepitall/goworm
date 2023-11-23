package worm

import (
	"github.com/creepitall/goworm/internal/models"
)

type Worm struct {
	positions models.Positions
}

func New() *Worm {
	return &Worm{positions: models.Positions{{X: 1, Y: 1}}}
}

func (w *Worm) Change(way models.Way) {
	head := w.positions[len(w.positions)-1]
	tail := w.positions[1:]

	out := make(models.Positions, 0)
	out = append(out, tail...)
	out = append(out, head.Change(way))

	w.positions = out
}

func (w *Worm) Add(way models.Way) {
	w.positions = append(w.positions, w.positions[len(w.positions)-1].Change(way))
}

func (w Worm) Positions() models.Positions {
	return w.positions
}
