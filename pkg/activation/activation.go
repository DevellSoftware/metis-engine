package activation

type Activation interface {
	Activate(float64) float64
	Derivative(float64) float64
}
