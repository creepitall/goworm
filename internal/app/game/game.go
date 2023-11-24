package game

import (
	"encoding/json"
	"fmt"

	"github.com/creepitall/goworm/internal/app/apple"
	"github.com/creepitall/goworm/internal/app/worm"

	"github.com/creepitall/goworm/internal/models"
)

type Worm interface {
	Add(way models.Way)
	Change(way models.Way)
	Positions() models.Positions
	GetHead() models.Position
}

type Apple interface {
	Add(x, y int)
	Positions() models.Positions
	Drop(p models.Positions)
	IsCrossed(p models.Position) bool
}

type Game struct {
	Worm  Worm
	Map   *Map
	Apple Apple
	Way   models.Way
	Death chan bool
}

func New() *Game {
	m := Map{}
	return &Game{
		Worm:  worm.New(),
		Map:   m.Init(640, 480),
		Apple: apple.New(),
		Death: make(chan bool),
		Way:   models.Right,
	}
}

func (g *Game) Conversion(way models.Way) []byte {
	g.Worm.Change(way)

	if g.Map.IsOutside(g.Worm.GetHead()) {
		fmt.Println("death")
		g.Death <- true
	}

	if g.Apple.IsCrossed(g.Worm.GetHead()) {
		g.Worm.Add(way)
	}

	g.Apple.Add(g.Map.Get())
	g.Apple.Drop(g.Worm.Positions())

	return g.toByte()
}

func (g Game) toByte() []byte {
	type Out struct {
		PositionPoint models.Positions `json:"positionPoint"`
		ApplePoint    models.Positions `json:"applePoint"`
	}

	out := Out{
		PositionPoint: g.Worm.Positions(),
		ApplePoint:    g.Apple.Positions(),
	}

	b, _ := json.Marshal(out)

	return b
}

func (g Game) IsDeath() chan bool {
	return g.Death
}

func (g *Game) ChangeWay(way models.Way) {
	g.Way = way
}

func (g *Game) GetWay() models.Way {
	return g.Way
}
