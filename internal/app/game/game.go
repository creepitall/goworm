package game

import (
	"encoding/json"

	"github.com/creepitall/goworm/internal/app/worm"

	"github.com/creepitall/goworm/internal/models"
)

type Worm interface {
	Change(way models.Way)
	Positions() models.Positions
}

type Game struct {
	Worm  Worm
	Way   models.Way
	Death chan bool
}

func New() *Game {
	return &Game{Worm: worm.New(), Death: make(chan bool), Way: models.Right}
}

func (g *Game) Conversion(way models.Way) []byte {
	g.Worm.Change(way)

	type Out struct {
		PositionPoint models.Positions `json:"positionPoint"`
	}

	out := Out{
		PositionPoint: g.Worm.Positions(),
	}

	b, _ := json.Marshal(out)

	return b
}

func (g *Game) ChangeWay(way models.Way) {
	g.Way = way
}

func (g *Game) GetWay() models.Way {
	return g.Way
}
