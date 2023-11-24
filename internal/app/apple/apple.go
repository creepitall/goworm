package apple

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/creepitall/goworm/internal/models"
)

type Apple struct {
	positions models.PositionsM
	r         *rand.Rand
}

func New() *Apple {
	return &Apple{
		positions: make(models.PositionsM, 0),
		r:         rand.New(rand.NewSource(99)),
	}
}

func (a *Apple) Add(x, y int) {
	if len(a.positions) == 3 {
		return
	}

	r := func(v int) int {
		return a.r.Intn(v)
	}

	xx := r(x)
	yy := r(y)
	xy := strings.Join([]string{strconv.Itoa(xx), strconv.Itoa(yy)}, "_")

	a.positions[xy] = models.Position{X: xx, Y: yy, XY: xy}
}

func (a *Apple) Drop(p models.Positions) {
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
