package game

import (
	"encoding/json"
	"fmt"

	"github.com/creepitall/goworm/internal/app/worm"

	"github.com/creepitall/goworm/internal/models"
)

type Worm interface {
	Change(way models.Way)
	Positions() models.Positions
	GetHead() models.Position
}

type Game struct {
	Worm  Worm
	Map   *Map
	Way   models.Way
	Death chan interface{}
}

func New() *Game {
	m := Map{}
	return &Game{
		Worm:  worm.New(),
		Map:   m.Init(640, 480),
		Death: make(chan interface{}),
		Way:   models.Right,
	}
}

func (g *Game) Conversion(way models.Way) []byte {
	g.Worm.Change(way)

	if g.Map.IsOutside(g.Worm.GetHead()) {
		fmt.Println("death")
		g.Death <- nil
	}

	type Out struct {
		PositionPoint models.Positions `json:"positionPoint"`
	}

	out := Out{
		PositionPoint: g.Worm.Positions(),
	}

	b, _ := json.Marshal(out)

	return b
}

func (g Game) IsDeath() chan interface{} {
	return g.Death
}

func (g *Game) ChangeWay(way models.Way) {
	g.Way = way
}

func (g *Game) GetWay() models.Way {
	return g.Way
}
