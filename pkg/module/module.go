package module

type Module interface {
	Forward()
	Backward(learningRate float64)
}
