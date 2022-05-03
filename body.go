package achilles

import "github.com/b5710546232/achilies/geom"

type Body struct {
	shape geom.Shape
	x     float64
	y     float64
}

func NewBody(shape geom.Shape, X, Y float64) *Body {
	o := &Body{
		shape: shape,
		x:     X,
		y:     Y,
	}
	return o
}

func (o *Body) GetPosition() geom.Vector2 {
	return o.shape.GetPosition()
}

func (o *Body) GetShape() geom.Shape {
	return o.shape
}

func (o *Body) Move(v geom.Vector2) {
	newX := o.shape.GetPosition().X + v.X
	newY := o.shape.GetPosition().Y + v.Y
	o.shape.UpdatePosition(newX, newY)
}

func (o *Body) getCircle() *geom.Circle {
	return o.shape.(*geom.Circle)
}

func (o *Body) Intersect(other *Body) bool {
	// o.GetShape().(type)
	switch o.GetShape().(type) {
	// self is circle
	case *geom.Circle:
		// other is circle
		switch o.GetShape().(type) {
		case *geom.Circle:
			return o.getCircle().IntersectOtherCircle(other.getCircle())
		}
	}
	return false
}
