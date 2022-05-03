package achilles

import (
	"time"
)

type System struct {
	world World
	Bodys []*Body
}

func NewSystem(w World) *System {
	return &System{
		world: w,
	}
}

func (s *System) AddBody(o *Body) {
	s.Bodys = append(s.Bodys, o)
}

func (s *System) Bodies() []*Body {
	return s.Bodys
}

// update loop
func (a *System) Run() {
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
