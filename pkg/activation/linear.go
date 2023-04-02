package activation

type Linear struct{}

func NewLinear() *Linear {
	return &Linear{}
}

func (l Linear) Activate(input float64) float64 {
	return input
}

func (l Linear) Derivative(input float64) float64 {
	return 1
}
