package life

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

// Используйте код из предыдущего урока по игре «Жизнь»
func NewWorld(height, width int) (*World, error) {
	var u World
	if height <= 0 || width <= 0 {
		return nil, fmt.Errorf("Wrong size")
	} else {
		u.Height = height
		u.Width = width
		u.Cells = make([][]bool, height)
		for i := range u.Cells {
			u.Cells[i] = make([]bool, width)
		}
	}
	return &u, nil
}

func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)
	return n == 3 || n == 2 && w.Alive(x, y)
}

func (w *World) neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && w.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (w *World) Alive(x, y int) bool {
	x = (x + w.Width) % w.Width
	y = (y + w.Height) % w.Height
	return w.Cells[y][x]
}

func NextState(oldWorld, newWorld *World) {
	for y := 0; y < oldWorld.Height; y++ {
		for x := 0; x < oldWorld.Width; x++ {
			newWorld.Cells[x][y] = oldWorld.next(x, y)
		}
	}
}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	// Количество живых клеток
	numAlive := percentage * w.Height * w.Width / 100
	// Заполним живыми первые клетки
	w.fillAlive(numAlive)
	// Получаем рандомные числа
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Рандомно меняем местами
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
