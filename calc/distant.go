package calc

import "math"

func Distant(p1x, p1y, p2x, p2y float64) float64 {
	return math.Sqrt(math.Pow(p2x-p1x, 2) + math.Pow(p2y-p1y, 2))
}
