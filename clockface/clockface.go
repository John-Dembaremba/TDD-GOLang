package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	return Point{X: 150, Y: 240}
}

// func zero() float64 {
// 	return 0.0
// }

func SecondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * (math.Pi / 30)
}

func SecondHandPoint(t time.Time) Point {
	r := SecondsInRadians(t)

	x := math.Sin((r))
	y := math.Cos(r)

	return Point{X: x, Y: y}
}
