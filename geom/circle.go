package geom

type Circle struct {
	Radius float64
	X      float64
	Y      float64
}

func NewCircle(x, y, radius float64) *Circle {
	c := &Circle{
		X:      x,
		Y:      y,
		Radius: radius,
	}
	return c
}

func (c *Circle) GetPosition() Vector2 {
	return Vector2{
		X: c.X,
		Y: c.Y,
	}
}

func (c *Circle) UpdatePosition(x, y float64) {
	c.X = x
	c.Y = y
}

func (c *Circle) Diameter() float64 {
	return c.Radius * 2
}

func (c *Circle) IntersectOtherCircle(other *Circle) bool {
	dist := c.GetPosition().Distant(other.GetPosition())
	return dist < c.Radius+other.Radius
}

func (c *Circle) GetRadius() float64 {
	return c.Radius * 2
}
