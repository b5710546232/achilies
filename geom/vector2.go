package geom

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) NewVector2(x, y float64) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

func (v Vector2) Distant(other Vector2) float64 {
	return math.Sqrt(math.Pow(other.X-v.X, 2) + math.Pow(other.Y-v.Y, 2))
}
