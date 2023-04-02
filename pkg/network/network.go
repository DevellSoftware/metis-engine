package network

import (
	"errors"

	"github.com/DevellSoftware/metis-engine/pkg/network/layer"
	"github.com/DevellSoftware/metis-engine/pkg/tensor"
)

type Network struct {
	layers []layer.Layer
}

func NewNetwork() *Network {
	return &Network{
		layers: make([]layer.Layer, 0),
	}
}

func (n *Network) Add(layer layer.Layer) {
	if len(n.layers) > 0 {
		n.layers[len(n.layers)-1].Connect(layer)
	}

	n.layers = append(n.layers, layer)
}

func (n *Network) Output() *tensor.Tensor {
	return n.layers[len(n.layers)-1].Output()
}

func (n *Network) Predict(input *tensor.Tensor) (*tensor.Tensor, error) {
	if len(n.layers) < 2 {
		return nil, errors.New("Network must have at least 2 layers")
	}
	n.layers[0].Set(input.Flip())
	n.layers[0].Forward()

	return n.layers[len(n.layers)-1].Output(), nil
}

func (n *Network) Train(input *tensor.Tensor, expectedOutput *tensor.Tensor, learningRate float64) {
	if len(n.layers) < 2 {
		return
	}

	n.layers[0].Set(input.Flip())
	n.layers[0].Forward()

	n.layers[len(n.layers)-1].SetTargetOutput(expectedOutput)
	n.layers[len(n.layers)-1].Backward(learningRate)
}
