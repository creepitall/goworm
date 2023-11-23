package app

import (
	"time"

	"github.com/creepitall/goworm/internal/app/game"

	"github.com/creepitall/goworm/internal/models"
)

type Runner struct {
	Game *game.Game
	t    *time.Ticker
	in   chan []byte
	out  chan []byte
}

func New() *Runner {
	g := game.New()

	return &Runner{Game: g,
		t:   time.NewTicker(2 * time.Second),
		in:  make(chan []byte, 100),
		out: make(chan []byte, 100),
	}
}

func (r *Runner) Run() {
	go func() {
		for _ = range r.t.C {
			r.Add([]byte(r.Game.GetWay()))
		}
	}()

	for {
		way := r.Game.GetWay()

		value, ok := <-r.in
		if ok {
			way = models.GetWayFromString(string(value))
			r.Game.ChangeWay(way)
		}

		r.out <- r.Game.Conversion(way)
	}
}

func (r *Runner) Add(b []byte) {
	r.in <- b
}

func (r *Runner) Get() chan []byte {
	return r.out
}
