package app

import (
	"fmt"
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
	//t := time.NewTicker(1 * time.Second)

	// defer func() { t.Stop() }()

	go func() {
		for _ = range r.t.C {
			r.Add([]byte(r.Game.GetWay()))
		}
	}()

	for {
		fmt.Println("were here")

		way := r.Game.GetWay()

		value, ok := <-r.in
		if ok {
			way = models.GetWayFromString(string(value))
			r.Game.ChangeWay(way)
		}

		r.out <- r.Game.Conversion(way)
	}

	// go func() {
	// 	for {
	// 		var way models.Way
	// 		select {
	// 		case value := <-r.in:
	// 			way = models.GetWayFromString(string(value))
	// 			r.Game.ChangeWay(way)
	// 		case <-t.C:
	// 			way = r.Game.GetWay()
	// 		}

	// 		r.out <- r.Game.Conversion(way)
	// 	}
	// }()
}

func (r *Runner) Add(b []byte) {
	r.in <- b
}

func (r *Runner) Get() chan []byte {
	return r.out
}
