package main

import (
	"image/color"

	achilles "github.com/b5710546232/achilies"
	"github.com/b5710546232/achilies/geom"
	"github.com/go-p5/p5"
)

type world struct {
	System       *achilles.System
	WorldObjects []*WorldObject
}

func (w *world) SetSystem(s *achilles.System) {
	w.System = s
}

func (w *world) AddObject(o *WorldObject) {
	w.System.AddObject(o.Body)
	w.WorldObjects = append(w.WorldObjects, o)
}

type WorldObject struct {
	Body  *achilles.Object
	Color color.RGBA
	IsHit bool
}

var system *achilles.System
var w *world
var circleObj *WorldObject
var circleObj2 *WorldObject
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
	if circleObj.Body.GetCircle().IsInterSectOtherCircle(circleObj2.Body.GetCircle()) {
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

func NewCircleObject(x, y, r float64) *WorldObject {
	circle := geom.NewCircle(x, y, r)
	o := achilles.NewObject(circle, x, y)
	obj := &WorldObject{
		Body: o,
	}

	obj.Color = color.RGBA{B: 255, A: 208} // blue
	return obj
}

func setup() {
	p5.Canvas(400, 400)
	p5.Background(color.Gray{Y: 220})
	circleObj = NewCircleObject(50, 50, 16)
	circleObj2 = NewCircleObject(150, 50, 48)

	// add object to world
	w.AddObject(circleObj2)
	w.AddObject(circleObj)
}

func draw() {
	p5.StrokeWidth(2)
	for _, obj := range w.WorldObjects {
		switch obj.Body.GetShape().(type) {
		case *geom.Circle:
			drawCircle(obj.Body.GetCircle(), obj.Color)
		}

	}

}

func drawCircle(circle *geom.Circle, c color.RGBA) {
	p5.Fill(c)
	p5.Circle(circle.X, circle.Y, circle.Diameter())
}
