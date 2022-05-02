package achilles

import (
	"time"
)

type System struct {
	world   World
	objects []*Object
}

func NewSystem(w World) *System {
	return &System{
		world: w,
	}
}

func (s *System) AddObject(o *Object) {
	s.objects = append(s.objects, o)
}

func (s *System) Objects() []*Object {
	return s.objects
}

// update loop
func (a *System) Update() {
	fps := time.Duration(time.Second / 60)
	ticker := time.NewTicker(fps)
	startTime := time.Now().UnixNano()
	go func() {
		for {
			<-ticker.C
			now := time.Now().UnixNano()
			delta := float64(now-startTime) / 1000000000
			startTime = now
			a.world.Update((delta))
		}
	}()
}
