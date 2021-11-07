package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type position struct {
	PositionName string `json:"position"`
}

type positionPoint struct {
	X           int         `json:"x"`
	Y           int         `json:"y"`
	MapSettings mapSettings `json:"mapSettings"`
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
	CurrentPosition     positionPoint
	CurrentMapSettings  mapSettings
	CurrentGameSettings gameSettings
	CurrentWay          string
	stopChan            chan bool = make(chan bool)
)

func main() {
	CurrentPosition = positionPoint{
		X: 1,
		Y: 1,
	}

	CurrentMapSettings = mapSettings{
		MaxX: 640,
		MaxY: 480,
		ObjX: 10,
		ObjY: 10,
	}

	CurrentWay = "right"

	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("web/templates/**/*")

	router.GET("/", getIndex)
	router.POST("/move", postMove)
	router.GET("/currentPosition", getCurrentPosition)
	router.POST("/currentWay", changeWay)
	//router.GET("/chunk", postChunk)
	router.POST("/changeGameSettings", postChangeGameSettings)

	router.Run("localhost:8080")
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
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
				CurrentPosition = move(CurrentWay)
			}
		}
	}()
}

func Stop() {
	stopChan <- true
}

func getCurrentPosition(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CurrentPosition)
}

func changeWay(c *gin.Context) {
	var newPosition position

	if err := c.BindJSON(&newPosition); err != nil {
		return
	}

	CurrentWay = newPosition.PositionName

	c.IndentedJSON(http.StatusOK, CurrentWay)
}

func postMove(c *gin.Context) {
	var newPosition position

	if err := c.BindJSON(&newPosition); err != nil {
		return
	}

	CurrentPosition = move(newPosition.PositionName)

	c.IndentedJSON(http.StatusOK, CurrentPosition)
}

func move(position string) positionPoint {
	curPositionX := CurrentPosition.X
	curPositionY := CurrentPosition.Y

	switch position {
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

	return positionPoint{
		X:           curPositionX,
		Y:           curPositionY,
		MapSettings: CurrentMapSettings,
	}
}
