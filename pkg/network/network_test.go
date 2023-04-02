package network

import (
	"testing"

	"github.com/DevellSoftware/metis-engine/pkg/activation"
	"github.com/DevellSoftware/metis-engine/pkg/network/layer"
	"github.com/DevellSoftware/metis-engine/pkg/tensor"
)

func TestFeedForward(t *testing.T) {
	n := NewNetwork()

	n.Add(layer.NewDense(2, activation.LinearFunction))
	n.Add(layer.NewDense(2, activation.ReLUFunction))
	n.Add(layer.NewDense(1, activation.ReLUFunction))

	epochs := 100
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

	result, error = n.Predict(tensor.NewTensor(tensor.FromArray([]float64{0, 0})))

	if error != nil {
		t.Error(error)
	}

	result.PrintDebug()

	for _, layer := range n.layers {
		layer.Output().PrintDebug()
		layer.Weights().PrintDebug()
	}
}
