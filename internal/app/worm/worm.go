package worm

import (
	"github.com/creepitall/goworm/internal/models"
)

type Worm struct {
	positions models.Positions
}

func New() *Worm {
	return &Worm{
		positions: models.Positions{{X: 1, Y: 1}},
	}
}

func (w *Worm) Change(way models.Way) {
	last := len(w.positions) - 1
	for i := 1; i < len(w.positions); i++ {
		w.positions[i-1] = w.positions[i]
	}
	w.positions[last] = w.positions[last].Change(way)

}

func (w Worm) IsCrossed() bool {
	uniq := make(map[models.Position]struct{})
	for _, k := range w.positions {
		if _, ok := uniq[k]; ok {
			return true
		}
		uniq[k] = struct{}{}
	}
	return false
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
