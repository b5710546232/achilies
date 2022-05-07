package collision

import (
	"github.com/b5710546232/achilies/calc"
)

func IntersectsRectPoint(rect Rect, point Point) bool {
	return point.GetX() >= rect.GetLeft() &&
		point.GetX() <= rect.GetRight() &&
		point.GetY() >= rect.GetTop() &&
		point.GetY() <= rect.GetBottom()

}

func IntersectsRectRect(a Rect, b Rect) bool {
	return !(a.GetRight() < b.GetLeft() || b.GetRight() < a.GetLeft() ||
		a.GetBottom() < b.GetTop() || b.GetBottom() < a.GetTop())
}

func Distant(p1 Point, p2 Point) float64 {
	return calc.Distant(p1.GetX(), p1.GetY(), p2.GetX(), p2.GetY())
}

func IntersectsCircleCircle(c1, c2 Circle) bool {
	dist := Distant(c1.GetPosition(), c2.GetPosition())
	return dist < c1.GetRadius()+c2.GetRadius()
}
