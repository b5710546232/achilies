package geom

type Shape interface {
	GetPosition() Vector2
	UpdatePosition(x, y float64)
}
