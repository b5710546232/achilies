package geom

import "github.com/b5710546232/achilies/calc"

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
	return calc.Distant(v.GetX(), v.GetY(), other.GetX(), other.GetY())
}

func (v Vector2) GetX() float64 {
	return v.X
}

func (v Vector2) GetY() float64 {
	return v.Y
}
