package main

import (
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
	X int `json:"x"`
	Y int `json:"y"`
}

type snakeData struct {
	PointData   []positionPoint `json:"positionPoint"`
	Length      int             `json:"length"`
	MapSettings mapSettings     `json:"mapSettings"`
}

type mapSettings struct {
	MaxX int `json:"maxX"`
	MaxY int `json:"maxY"`
	// Размер объекта (квадратика)
	ObjX int `json:"objX"`
	ObjY int `json:"objY"`
}

type gameSettings struct {
	GameStart bool `json:"gameStart"`
	GameReset bool `json:"gameReset"`
}

var (
	CurrentSnake        snakeData
	CurrentMapSettings  mapSettings
	CurrentGameSettings gameSettings
	CurrentWay          string
	NewWay              string
	stopChan            chan bool = make(chan bool)
)

func main() {
	var curPos []positionPoint
	curPos = append(curPos, positionPoint{
		X: 0,
		Y: 0,
	})

	CurrentMapSettings = mapSettings{
		MaxX: 640,
		MaxY: 480,
		ObjX: 40,
		ObjY: 40,
	}

	CurrentSnake = snakeData{
		PointData:   curPos,
		Length:      1,
		MapSettings: CurrentMapSettings,
	}

	CurrentWay = "right"

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/currentPosition", getCurrentPosition)
	router.GET("/addTail", postAddTail)
	router.POST("/currentWay", changeWay)
	//router.GET("/chunk", postChunk)
	router.POST("/changeGameSettings", postChangeGameSettings)

	router.Run("localhost:8080")
}

func postAddTail(c *gin.Context) {
	go CurrentSnake.changeLength()

	c.IndentedJSON(http.StatusOK, CurrentSnake)
}

func postChangeGameSettings(c *gin.Context) {
	var reqGameSettings gameSettings
	if err := c.BindJSON(&reqGameSettings); err != nil {
		return
	}

	if !CurrentGameSettings.GameStart && reqGameSettings.GameStart {
		go Run()
	}

	if !reqGameSettings.GameStart {
		Stop()
	}

	CurrentGameSettings.GameStart = reqGameSettings.GameStart

	c.IndentedJSON(http.StatusOK, CurrentGameSettings)
}

func Run() {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-stopChan:
				ticker.Stop()
				return
			case <-ticker.C:
				CurrentSnake.changePosition()
			}
		}
	}()
}

func Stop() {
	stopChan <- true
}

func getCurrentPosition(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CurrentSnake)
}

func changeWay(c *gin.Context) {
	var newPosition position

	if err := c.BindJSON(&newPosition); err != nil {
		return
	}

	NewWay = newPosition.SideName

	c.IndentedJSON(http.StatusOK, CurrentWay)
}

func (s *snakeData) changePosition() {
	var wayName string
	var newPointData []positionPoint
	for _, value := range s.PointData {
		wayName = value.changePoint(CurrentWay, NewWay)
		newPointData = append(newPointData, value)
	}
	s.PointData = newPointData

	CurrentWay = wayName
}

func (s *snakeData) changeLength() {
	currLen := s.Length
	s.Length++

	headX := s.PointData[currLen-1].X
	headY := s.PointData[currLen-1].Y

	var newPoint positionPoint
	newPoint.addTail(headX, headY)
	s.PointData = append(s.PointData, newPoint)
}

func (p *positionPoint) addTail(headX, headY int) {
	switch CurrentWay {
	case "right":
		p.X = headX - 1
		p.Y = headY
	case "left":
		p.X = headX + 1
		p.Y = headY
	case "up":
		p.X = headX
		p.Y = headY + 1
	case "down":
		p.X = headX
		p.Y = headY - 1

	default:
		p.X = headX
		p.Y = headY
	}
}

func (p *positionPoint) changePoint(currentSideName, newSideName string) (wayName string) {
	curPositionX := p.X
	curPositionY := p.Y

	if newSideName == "" {
		newSideName = currentSideName
	}

	switch newSideName {
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

	return
}
