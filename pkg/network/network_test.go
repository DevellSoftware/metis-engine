package network

import (
	"testing"

	"github.com/DevellSoftware/metis/pkg/activation"
	"github.com/DevellSoftware/metis/pkg/log"
	"github.com/DevellSoftware/metis/pkg/network/layer"
	"github.com/DevellSoftware/metis/pkg/tensor"
)

func TestFeedForward(t *testing.T) {
	n := NewNetwork()

	n.Add(layer.NewDense(2, activation.LinearFunction))
	n.Add(layer.NewDense(2, activation.SigmoidFunction))
	n.Add(layer.NewDense(1, activation.SigmoidFunction))

	epochs := 30
	lr := 0.1

	for i := 0; i < epochs; i++ {
		n.Train(tensor.NewTensor(tensor.FromArray([]float64{0, 0})), tensor.NewTensor(tensor.FromArray([]float64{0})), lr)
		n.Train(tensor.NewTensor(tensor.FromArray([]float64{0, 1})), tensor.NewTensor(tensor.FromArray([]float64{1})), lr)
		n.Train(tensor.NewTensor(tensor.FromArray([]float64{1, 0})), tensor.NewTensor(tensor.FromArray([]float64{1})), lr)
		n.Train(tensor.NewTensor(tensor.FromArray([]float64{1, 1})), tensor.NewTensor(tensor.FromArray([]float64{0})), lr)
	}

	result, error := n.Predict(tensor.NewTensor(tensor.FromArray([]float64{0, 1})))

	if error != nil {
		t.Error(error)
	}

	log.Log("Result")
	result.PrintDebug()

	result, error = n.Predict(tensor.NewTensor(tensor.FromArray([]float64{0, 0})))

	if error != nil {
		t.Error(error)
	}

	log.Log("Result")
	result.PrintDebug()
}
