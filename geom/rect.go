package geom

import "github.com/b5710546232/achilies/collision"

type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (r *Rect) GetLeft() float64 {
	return r.X - r.W/2
}

func (r *Rect) GetRight() float64 {
	return r.X + r.W/2
}

func (r *Rect) GetTop() float64 {
	return r.Y - r.H/2
}

func (r *Rect) GetX() float64 {
	return r.X
}

func (r *Rect) GetY() float64 {
	return r.Y
}

func (r *Rect) GetBottom() float64 {
	return r.Y + r.H/2
}

func (r *Rect) GetBottomLeft() Point {
	return Point{
		X: r.GetLeft(),
		Y: r.GetBottom(),
	}
}

func (r *Rect) UpdatePosition(x, y float64) {
	r.X = x
	r.Y = y
}

func (r *Rect) GetTopRight() Point {
	return Point{
		X: r.GetRight(),
		Y: r.GetTop(),
	}
}

func (r *Rect) ContainsPoint(point *Point) bool {
	return collision.IntersectsRectPoint(r, point)
}

func (r *Rect) IntersectsRect(other *Rect) bool {
	return collision.IntersectsRectRect(r, other)
}
