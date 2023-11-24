package app

import (
	"time"

	"github.com/creepitall/goworm/internal/app/game"

	"github.com/creepitall/goworm/internal/models"
)

type Runner struct {
	Game *game.Game
	in   chan []byte
	out  chan []byte
}

func New() *Runner {
	g := game.New()

	return &Runner{Game: g,
		in:  make(chan []byte, 100),
		out: make(chan []byte, 100),
	}
}

func (r *Runner) Run() {
	t := time.NewTicker(500 * time.Millisecond)

	defer func() {
		t.Stop()
		close(r.in)
		close(r.out)
	}()

	for {
		select {
		case <-r.Game.IsDeath():
			return
		case value := <-r.in:
			way := models.GetWayFromString(string(value))
			if !r.Game.GetWay().IsCrossed(way) {
				r.Game.ChangeWay(way)
			}
		case <-t.C:
			way := r.Game.GetWay()
			r.out <- r.Game.Conversion(way)
		}
	}
}

func (r *Runner) Add(b []byte) {
	r.in <- b
}

func (r *Runner) Get() chan []byte {
	return r.out
}
