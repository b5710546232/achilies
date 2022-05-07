package collision

type Rect interface {
	GetLeft() float64
	GetRight() float64
	GetTop() float64
	GetBottom() float64
}

type Point interface {
	GetX() float64
	GetY() float64
}

type Circle interface {
	GetPosition() Point
	GetRadius() float64
}
