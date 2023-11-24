package apple

import (
	"math/rand"

	"github.com/creepitall/goworm/internal/models"
)

type Area interface {
	Get() (int, int)
}

type Apple struct {
	positions models.PositionsM
	r         *rand.Rand
	area      Area
	limit     int
}

func New(area Area) *Apple {
	return &Apple{
		positions: make(models.PositionsM, 0),
		r:         rand.New(rand.NewSource(99)),
		area:      area,
		limit:     3,
	}
}

func (a *Apple) Add(p models.Positions) {
	a.add()
	a.drop(p)
}

func (a *Apple) add() {
	if len(a.positions) == a.limit {
		return
	}

	r := func(v int) int {
		return a.r.Intn(v)
	}

	x, y := a.area.Get()

	pos := models.Position{}.Fill(r(x), r(y))

	a.positions[pos.XY] = pos
}

func (a *Apple) drop(p models.Positions) {
	toDrop := make([]string, 0, len(p))
	for _, v := range p {
		if _, ok := a.positions[v.XY]; ok {
			toDrop = append(toDrop, v.XY)
		}
	}

	for _, v := range toDrop {
		delete(a.positions, v)
	}
}

func (a Apple) Positions() models.Positions {
	pos := make(models.Positions, 0, len(a.positions))
	for _, v := range a.positions {
		pos = append(pos, v)
	}
	return pos
}

func (a Apple) IsCrossed(p models.Position) bool {
	if _, ok := a.positions[p.XY]; ok {
		delete(a.positions, p.XY)
		return true
	}
	return false
}
