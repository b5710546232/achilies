package main

import (
	"image/color"

	achilles "github.com/b5710546232/achilies"
	"github.com/b5710546232/achilies/geom"
	"github.com/go-p5/p5"
)

type world struct {
	System  *achilles.System
	Objects []*Object
}

func (w *world) SetSystem(s *achilles.System) {
	w.System = s
}

func (w *world) AddBody(o *Object) {
	w.System.AddBody(o.Body)
	w.Objects = append(w.Objects, o)
}

type Object struct {
	Body     *achilles.Body
	Color    color.RGBA
	IsHit    bool
	Diameter float64
}

var system *achilles.System
var w *world
var circleObj *Object
var circleObj2 *Object
var speedX = float64(100)
var dirX = 1.0

func (u *world) Update(dt float64) {
	if circleObj.Body.GetPosition().X > 300 {
		dirX = -1.0
	}
	if circleObj.Body.GetPosition().X < 0 {
		dirX = 1.0
	}

	v := geom.Vector2{
		X: speedX * dirX * dt,
		Y: 0,
	}
	circleObj.Body.Move(v)
	if circleObj.Body.Intersect(circleObj2.Body) {
		circleObj.IsHit = true
		circleObj.Color = color.RGBA{R: 255, A: 208}
	} else {
		circleObj.IsHit = false
		circleObj.Color = color.RGBA{B: 255, A: 208}
	}
}

func main() {
	w = &world{}
	system = achilles.NewSystem(w)
	w.SetSystem(system)

	system.Run()
	p5.Run(setup, draw)
}

func NewCircleBody(x, y, r float64) *Object {
	circle := geom.NewCircle(x, y, r)
	o := achilles.NewBody(circle, x, y)
	obj := &Object{
		Body:     o,
		Diameter: r * 2,
	}

	obj.Color = color.RGBA{B: 255, A: 208} // blue
	return obj
}

func setup() {
	p5.Canvas(400, 400)
	p5.Background(color.Gray{Y: 220})
	circleObj = NewCircleBody(50, 50, 16)
	circleObj2 = NewCircleBody(150, 50, 48)

	// add Body to world
	w.AddBody(circleObj2)
	w.AddBody(circleObj)
}

func draw() {
	p5.StrokeWidth(2)
	for _, obj := range w.Objects {
		switch obj.Body.GetShape().(type) {
		case *geom.Circle:
			drawCircle(obj, obj.Color)
		}

	}

}

func drawCircle(obj *Object, c color.RGBA) {
	p5.Fill(c)
	pos := obj.Body.GetPosition()
	p5.Circle(pos.X, pos.Y, obj.Diameter)
}
