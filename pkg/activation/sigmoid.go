package activation

import "math"

type Sigmoid struct{}

func (s Sigmoid) Activate(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func (s Sigmoid) Derivative(x float64) float64 {
	return x * (1 - x)
}
