package main

import (
	"image/color"

	achilles "github.com/b5710546232/achilies"
	"github.com/b5710546232/achilies/geom"
	"github.com/go-p5/p5"
)

type world struct{}

var system *achilles.System
var circleObj *achilles.Object
var circleObj2 *achilles.Object
var speedX = float64(100)
var isHitToggle = false
var dirX = 1.0

func (u *world) Update(dt float64) {
	if circleObj.GetPosition().X > 300 {
		dirX = -1.0
	}
	if circleObj.GetPosition().X < 0 {
		dirX = 1.0
	}

	v := geom.Vector2{
		X: speedX * dirX * dt,
		Y: 0,
	}
	circleObj.Move(v)
	if circleObj.GetCircle().IsInterSectOtherCircle(circleObj2.GetCircle()) {
		if !isHitToggle {
			isHitToggle = true
		}
	} else {
		isHitToggle = false
	}
}

func main() {
	w := &world{}
	system = achilles.NewSystem(w)
	system.Update()
	p5.Run(setup, draw)
}

func NewCircleObject(x, y, r float64) *achilles.Object {
	circle := geom.NewCircle(x, y, r)
	o := achilles.NewObject(circle, x, y)
	return o
}

func setup() {
	p5.Canvas(400, 400)
	p5.Background(color.Gray{Y: 220})
	circleObj = NewCircleObject(50, 50, 16)
	circleObj2 = NewCircleObject(150, 50, 48)
	system.AddObject(circleObj)
	system.AddObject(circleObj2)
}

func draw() {
	p5.StrokeWidth(2)
	p5.Fill(color.RGBA{R: 255, A: 208})
	p5.Circle(circleObj.GetCircle().X, circleObj.GetCircle().Y, circleObj.GetCircle().Diameter())

	if isHitToggle {
		p5.Fill(color.RGBA{B: 255, A: 208})
		p5.Circle(circleObj2.GetCircle().X, circleObj2.GetCircle().Y, circleObj2.GetCircle().Diameter())

	} else {
		p5.Fill(color.RGBA{G: 255, A: 208})
		p5.Circle(circleObj2.GetCircle().X, circleObj2.GetCircle().Y, circleObj2.GetCircle().Diameter())
	}

}
