package layer

import "github.com/DevellSoftware/metis-engine/pkg/tensor"

type Layer interface {
	Output() *tensor.Tensor
	Connect(layer Layer)
	ConnectInput(layer Layer)
	Forward()
	Backward(learningRate float64)
	SetTargetOutput(targetOutput *tensor.Tensor)
	Set(input *tensor.Tensor)
	Error() *tensor.Tensor
	Size() int
	Weights() *tensor.Tensor
}
