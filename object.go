package achilles

import "github.com/b5710546232/achilies/geom"

type Object struct {
	shape geom.Shape
	x     float64
	y     float64
}

func NewObject(shape geom.Shape, X, Y float64) *Object {
	o := &Object{
		shape: shape,
		x:     X,
		y:     Y,
	}
	return o
}

func (o Object) GetCircle() *geom.Circle {
	return o.shape.(*geom.Circle)
}

func (o Object) GetPosition() geom.Vector2 {
	return o.shape.GetPosition()
}

func (o Object) GetShape() geom.Shape {
	return o.shape
}

func (o Object) Move(v geom.Vector2) {
	newX := o.shape.GetPosition().X + v.X
	newY := o.shape.GetPosition().Y + v.Y
	o.shape.UpdatePosition(newX, newY)
}
