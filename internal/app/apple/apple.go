package apple

import (
	"github.com/creepitall/goworm/internal/models"
)

const limit = 3

type Apple struct {
	positions models.Positions
}

func New() *Apple {
	return &Apple{
		positions: make(models.Positions, 0, limit),
	}
}

func (a Apple) IsCanAdd() bool {
	return len(a.positions) != limit
}

func (a *Apple) Add(p models.Position) {
	a.positions = append(a.positions, p)
}

func (a Apple) Positions() models.Positions {
	return a.positions
}

func (a *Apple) IsCrossed(p models.Position) bool {
	for i, v := range a.positions {
		if v == p {
			switch i {
			case 0:
				a.positions = a.positions[i+1:]
			case len(a.positions) - 1:
				a.positions = a.positions[:i]
			default:
				new := a.positions[:i]
				new = append(new, a.positions[i+1:]...)
				a.positions = new
			}
			return true
		}
	}
	return false
}
