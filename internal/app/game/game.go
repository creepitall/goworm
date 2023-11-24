package game

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/creepitall/goworm/internal/app/apple"
	"github.com/creepitall/goworm/internal/app/area"
	"github.com/creepitall/goworm/internal/app/worm"

	"github.com/creepitall/goworm/internal/models"
)

type Worm interface {
	Add(way models.Way)
	Change(way models.Way) bool
	Positions() models.Positions
	GetHead() models.Position
}

type Apple interface {
	Add(models.Positions)
	Positions() models.Positions
	IsCrossed(p models.Position) bool
}

type Game struct {
	Worm  Worm
	Apple Apple
	Way   models.Way
	Death chan bool
	mu    sync.RWMutex
}

func New() *Game {
	area := area.New(640, 480)
	return &Game{
		Worm:  worm.New(area),
		Apple: apple.New(area),
		Death: make(chan bool),
		Way:   models.Right,
	}
}

func (g *Game) Conversion() []byte {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.Worm.Change(g.Way) {
		fmt.Println("death")
		g.Death <- true
	}

	if g.Apple.IsCrossed(g.Worm.GetHead()) {
		g.Worm.Add(g.Way)
	}

	g.Apple.Add(g.Worm.Positions())

	return g.toByte()
}

func (g *Game) toByte() []byte {
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

func (g *Game) IsDeath() chan bool {
	return g.Death
}

func (g *Game) ChangeWay(way models.Way) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.Way.IsCrossed(way) {
		g.Way = way
	}
}
