package area

import (
	"math/rand"

	"github.com/creepitall/goworm/internal/models"
)

type Area struct {
	width  int
	height int
	X      int
	Y      int
	r      *rand.Rand
}

func New(w, h int) Area {
	return Area{
		width:  w,
		height: h,
		X:      w / 40,
		Y:      h / 40,
		r:      rand.New(rand.NewSource(99)),
	}
}

func (a Area) IsOutside(p models.Position) bool {
	x, y := p.Get()
	if x < 0 || y < 0 {
		return true
	}
	if x >= a.X || y >= a.Y {
		return true
	}
	return false
}

func (a Area) MakeApple(apples, worm models.Positions) models.Position {
	r := func(m map[int]struct{}, v int) int {
		nums := make([]int, 0)
		for i := 1; i <= v; i++ {
			if _, find := m[i]; find {
				continue
			}
			nums = append(nums, i)
		}
		idx := a.r.Intn(len(nums))
		return nums[idx]
	}

	memo := make(map[int]struct{})
	var i int
	for i < len(apples) && i < len(worm) {
		if i < len(apples) {
			v := apples[i]
			memo[v.X] = struct{}{}
			memo[v.Y] = struct{}{}
		}
		if i < len(worm) {
			v := worm[i]
			memo[v.X] = struct{}{}
			memo[v.Y] = struct{}{}
		}
		i++
	}
	x, y := a.Get()

	return models.Position{X: r(memo, x), Y: r(memo, y)}
}

func (a Area) Get() (int, int) {
	return a.X, a.Y
}
