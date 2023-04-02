package activation

import "math"

type ReLU struct{}

func (r ReLU) Activate(x float64) float64 {
	return math.Max(x, 0)
}

func (r ReLU) Derivative(x float64) float64 {
	if x < 0 {
		return 0
	}
	return 1
}
