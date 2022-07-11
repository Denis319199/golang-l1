package task24

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func Distance(f, s Point) float64 {
	diffX := s.X - f.X
	diffY := s.Y - f.Y

	return math.Sqrt(diffX*diffX + diffY*diffY)
}

func Task24() {
	fmt.Println(Distance(Point{-1, 3}, Point{6, 2}))
}
