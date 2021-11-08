package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type position struct {
	SideName string `json:"position"`
}

// TODO way name
type positionPoint struct {
	X        int    `json:"x"`
	Y        int    `json:"y"`
	SideName string `json:"sideName"`
}

type snakeData struct {
	PointData []positionPoint `json:"positionPoint"`
	Length    int             `json:"length"`
	Death     bool            `json:"Death"`
	Chunk     []positionPoint `json:"chunkPoint"`
}

type mapSettings struct {
	MaxX int `json:"maxX"`
	MaxY int `json:"maxY"`
	// Размер объекта (квадратика)
	ObjX int `json:"objX"`
	ObjY int `json:"objY"`
}

type gameSettings struct {
	GameStart   bool        `json:"gameStart"`
	GameReset   bool        `json:"gameReset"`
	ChunkToDeth int         `json:"chunkToDeth"`
	MapSettings mapSettings `json:"mapSettings"`
}

var (
	CurrentSnake        snakeData
	CurrentMapSettings  mapSettings
	CurrentGameSettings gameSettings
	CurrentWay          string
)

func main() {
	initSettings()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/currentPosition", getCurrentPosition)
	router.GET("/requestChunk", getRequestChunk)
	router.POST("/currentWay", changeWay)
	router.POST("/changeGameSettings", postChangeGameSettings)

	router.Run("localhost:8080")
}

func getRequestChunk(c *gin.Context) {
	//go CurrentSnake.changeLength()
	CurrentSnake.getChunk()

	c.IndentedJSON(http.StatusOK, CurrentSnake)
}

func (s *snakeData) getChunk() {
	mx := CurrentMapSettings.MaxX/CurrentMapSettings.ObjX - 1
	my := CurrentMapSettings.MaxY/CurrentMapSettings.ObjY - 1

	var exit bool
	var pp positionPoint
	for {
		x, y := Shuffle(mx, my)
		pp = positionPoint{
			X: x,
			Y: y,
		}
		for _, value := range s.PointData {
			if value.X != pp.X && value.Y != pp.Y {
				exit = true
				break
			}
		}
		if exit {
			break
		}
	}

	s.Chunk = append(s.Chunk, positionPoint{
		X: pp.X,
		Y: pp.Y,
	})
}

func Shuffle(maxx, maxy int) (x, y int) {
	rand.Seed(time.Now().UTC().UnixNano())
	vlX := randInt(1, maxx)
	vlY := randInt(1, maxy)

	return vlX, vlY
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func initSettings() {
	var curPos []positionPoint
	curPos = append(curPos, positionPoint{
		X:        0,
		Y:        0,
		SideName: "right",
	})

	CurrentMapSettings = mapSettings{
		MaxX: 640,
		MaxY: 480,
		ObjX: 40,
		ObjY: 40,
	}

	CurrentGameSettings = gameSettings{
		GameStart:   false,
		GameReset:   false,
		ChunkToDeth: 5,
		MapSettings: CurrentMapSettings,
	}

	CurrentSnake = snakeData{
		PointData: curPos,
		Length:    1,
		Death:     false,
		Chunk:     make([]positionPoint, 0),
	}

	CurrentWay = "right"
}

func postChangeGameSettings(c *gin.Context) {
	var reqGameSettings gameSettings
	if err := c.BindJSON(&reqGameSettings); err != nil {
		return
	}

	CurrentGameSettings.GameStart = reqGameSettings.GameStart

	if reqGameSettings.GameReset {
		initSettings()
	}

	c.IndentedJSON(http.StatusOK, CurrentGameSettings)
}

func getCurrentPosition(c *gin.Context) {
	if CurrentGameSettings.GameStart {
		CurrentSnake.changePosition()

		CurrentSnake.actualStatus()
		CurrentGameSettings.GameStart = !CurrentSnake.Death
	}
	c.IndentedJSON(http.StatusOK, CurrentSnake)
}

func changeWay(c *gin.Context) {
	var newPosition position

	if err := c.BindJSON(&newPosition); err != nil {
		return
	}

	CurrentWay = newPosition.SideName

	c.IndentedJSON(http.StatusOK, CurrentWay)
}

func (s *snakeData) actualStatus() {
	var headX, headY int

	if len(s.Chunk) > CurrentGameSettings.ChunkToDeth {
		s.Death = true
	}

	for id, value := range s.PointData {
		if id == 0 {
			headX = value.X
			headY = value.Y
			for idCh, ch := range s.Chunk {
				if ch.X == headX && ch.Y == headY {
					go s.changeLength()
					s.Chunk[idCh] = s.Chunk[len(s.Chunk)-1]
					s.Chunk = s.Chunk[:len(s.Chunk)-1]
					break
				}
			}
		} else {
			if headX == value.X && headY == value.Y {
				s.Death = true
				break
			}
		}
	}
}

func (s *snakeData) changePosition() {
	var newPointData []positionPoint
	isFirst := true
	var a, b positionPoint
	for _, value := range s.PointData {
		a, b = b, value
		if isFirst {
			value.changePoint()

			newPointData = append(newPointData, value)
			isFirst = false
		} else {
			newPointData = append(newPointData, a)
		}
	}

	s.PointData = newPointData
}

func (s *snakeData) changeLength() {
	currLen := s.Length
	s.Length++

	tailX := s.PointData[currLen-1].X
	tailY := s.PointData[currLen-1].Y
	tailSide := s.PointData[currLen-1].SideName

	var newPoint positionPoint
	newPoint.addTail(tailX, tailY, tailSide)
	s.PointData = append(s.PointData, newPoint)
}

func (p *positionPoint) addTail(tailX, tailY int, tailSideName string) {
	switch tailSideName {
	case "right":
		p.X = tailX - 1
		p.Y = tailY
	case "left":
		p.X = tailX + 1
		p.Y = tailY
	case "up":
		p.X = tailX
		p.Y = tailY + 1
	case "down":
		p.X = tailX
		p.Y = tailY - 1

	default:
		p.X = tailX
		p.Y = tailY
	}
	p.SideName = tailSideName
}

func (p *positionPoint) changePoint() {
	curPositionX := p.X
	curPositionY := p.Y

	switch CurrentWay {
	case "right":
		if (curPositionX+1)*CurrentMapSettings.ObjX < CurrentMapSettings.MaxX {
			curPositionX = curPositionX + 1
		}
	case "left":
		if curPositionX-1 > 0 {
			curPositionX = curPositionX - 1
		} else {
			curPositionX = 0
		}
	case "up":
		if curPositionY-1 > 0 {
			curPositionY = curPositionY - 1
		} else {
			curPositionY = 0
		}
	case "down":
		if (curPositionY+1)*CurrentMapSettings.ObjY < CurrentMapSettings.MaxY {
			curPositionY = curPositionY + 1
		}
	default:
		curPositionX = 0
		curPositionY = 0
	}

	p.X = curPositionX
	p.Y = curPositionY
	p.SideName = CurrentWay
}
