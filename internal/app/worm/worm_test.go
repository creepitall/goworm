package worm

import (
	"fmt"
	"testing"

	"github.com/creepitall/goworm/internal/models"
)

func TestChange(t *testing.T) {

	t.Run("", func(t *testing.T) {
		worm := New()

		worm.Change(models.Right) // {1 1}
		fmt.Println(worm.Positions())
		worm.Change(models.Right) // {1 1}
		fmt.Println(worm.Positions())
		worm.Change(models.Right) // {1 1}
		fmt.Println(worm.Positions())
		worm.Change(models.Right) // {1 1}
		fmt.Println(worm.Positions())
		worm.Add(models.Right) // {3 1} {3 1}
		fmt.Println(worm.Positions())
		worm.Change(models.Right) // {4 1} {5 1}
		fmt.Println(worm.Positions())
		worm.Change(models.Right) // {5 1} {7 1}
		fmt.Println(worm.Positions())
		worm.Change(models.Down)
		fmt.Println(worm.Positions())
		worm.Change(models.Down)
		fmt.Println(worm.Positions())
	})
}

var _pos = models.Positions{
	{X: 1, Y: 1},
}
