package main

import(
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func move_cursor_to(x int, y int) string {
	return "" + "[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H"
}

type Star struct {
	x int
	y int
	init int
	speed float64
	shape string
}

// see http://ascii-table.com/ansi-escape-sequences-vt-100.php
func (s *Star) render(i int) string {
	return move_cursor_to(s.x, s.y) +
		"[38;5;" + strconv.Itoa(232 + (s.init + int(math.Abs(float64(i) * s.speed))) % 24) + "m" +
		s.shape +
		"[39m"
}

func main() {
	rand.Seed(42)
	stars := []Star{}

	for i := 0; i < 256; i++ {
		str := "."
		if rand.Intn(10) == 9 {
			str = "*"
		}
		stars = append(stars, Star{rand.Intn(200), rand.Intn(60), rand.Intn(23), 1, str})
	}

	for {
		for i := -23; i < 23; i++ {
			time.Sleep(42 * time.Millisecond)
			str := "[2J"
			for j := 0; j < len(stars); j++ {
				str += stars[j].render(i)
			}
			fmt.Print(str + move_cursor_to(0, 0))
		}
	}
}
