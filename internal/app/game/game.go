package game

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/creepitall/goworm/internal/app/apple"
	"github.com/creepitall/goworm/internal/app/area"
	"github.com/creepitall/goworm/internal/app/worm"

	"github.com/creepitall/goworm/internal/models"
)

type areaiface interface {
	IsOutside(worm models.Position) bool
	MakeApple(apples, worm models.Positions) models.Position
}

type wormiface interface {
	Add(way models.Way)
	Change(way models.Way)
	Positions() models.Positions
	GetHead() models.Position
	IsCrossed() bool
}

type appleiface interface {
	Add(p models.Position)
	Positions() models.Positions
	IsCrossed(worm models.Position) bool
	IsCanAdd() bool
}

type Game struct {
	area  areaiface
	worm  wormiface
	apple appleiface
	Way   models.Way
	Exit  chan bool
	mu    sync.RWMutex
}

func New() *Game {
	area := area.New(640, 480)
	return &Game{
		area:  area,
		worm:  worm.New(),
		apple: apple.New(),
		Exit:  make(chan bool),
		Way:   models.Stop,
	}
}

func (g *Game) Conversion() []byte {
	g.mu.RLock()
	defer g.mu.RUnlock()

	log.Print(g.worm.Positions())

	g.worm.Change(g.Way)
	switch {
	case g.area.IsOutside(g.worm.GetHead()), g.worm.IsCrossed():
		fmt.Println("death")
		g.Exit <- true
	}
	if g.apple.IsCrossed(g.worm.GetHead()) {
		g.worm.Add(g.Way)
	}
	if g.apple.IsCanAdd() {
		apple := g.area.MakeApple(g.apple.Positions(), g.worm.Positions())
		g.apple.Add(apple)
	}

	return g.toByte()
}

func (g *Game) IsExit() chan bool {
	return g.Exit
}

func (g *Game) ChangeWay(way models.Way) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.Way.IsCrossed(way) {
		g.Way = way
	}
}

func (g *Game) toByte() []byte {
	type Out struct {
		PositionPoint models.Positions `json:"positionPoint"`
		ApplePoint    models.Positions `json:"applePoint"`
	}

	out := Out{
		PositionPoint: g.worm.Positions(),
		ApplePoint:    g.apple.Positions(),
	}

	b, _ := json.Marshal(out)

	return b
}
