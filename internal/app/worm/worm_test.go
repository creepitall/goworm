package worm

import (
	"fmt"
	"testing"

	"github.com/creepitall/goworm/internal/models"
)

func TestMyFunc(t *testing.T) {

	t.Run("change position", func(t *testing.T) {
		w := New()
		for i := 0; i < 5; i++ {
			if i == 3 {
				w.Add(models.Right)
			}
			w.Change(models.Right)
			fmt.Println(w.Positions())
		}
	})

}

func BenchmarkSample(b *testing.B) {
	b.Run("change position old", func(b *testing.B) {
		w := New()
		for i := 0; i < 5; i++ {
			if i == 3 {
				w.Add(models.Right)
			}
			w.Change(models.Right)
		}
	})
}
