package layer

import (
	"math/rand"
	"time"

	"github.com/DevellSoftware/metis/pkg/activation"
	"github.com/DevellSoftware/metis/pkg/log"
	"github.com/DevellSoftware/metis/pkg/tensor"
)

type Dense struct {
	value        *tensor.Tensor
	weights      *tensor.Tensor
	inputValue   *tensor.Tensor
	targetOutput *tensor.Tensor
	size         int

	output Layer
	input  Layer

	activation activation.ActivationType
}

func NewDense(size int, activationFunction activation.ActivationType) *Dense {
	return &Dense{
		size:       size,
		weights:    tensor.NewTensor(tensor.WithRank(2)),
		value:      tensor.NewTensor(tensor.WithRank(1)),
		inputValue: tensor.NewTensor(tensor.WithRank(2)),
		activation: activationFunction,
	}
}

func (d *Dense) Output() *tensor.Tensor {
	return d.value.Activate(d.activation)
}

func (d *Dense) ConnectInput(layer Layer) {
	d.input = layer
	d.initializeWeights()
}

func (d *Dense) initializeWeights() {
	rand.Seed(time.Now().UnixMicro())

	if d.input != nil {
		for neuronIndex := 0; neuronIndex < d.size; neuronIndex++ {
			for neuronInputIndex := 0; neuronInputIndex < d.input.Size(); neuronInputIndex++ {
				d.weights.Set(
					tensor.NewIndex(neuronInputIndex, neuronIndex),
					2*rand.Float64()-1,
				)

				log.Log("Initialized weights for layer with size %d", d.size)
				d.weights.PrintDebug()
			}
		}
	}
}

func (d *Dense) Weights() *tensor.Tensor {
	return d.weights
}

func (d *Dense) Connect(layer Layer) {
	d.output = layer
	layer.ConnectInput(d)

}

func (d *Dense) Forward() {
	if d.input == nil {
		d.value = d.inputValue
	} else {
		d.value = d.weights.Multiply(d.inputValue)
	}

	if d.output != nil {
		d.output.Set(d.Output())

		d.output.Forward()
	}
}

func (d *Dense) SetTargetOutput(targetOutput *tensor.Tensor) {
	d.targetOutput = targetOutput
}

func (d *Dense) Size() int {
	return d.size
}

func (d *Dense) Error() *tensor.Tensor {
	if d.output == nil && d.targetOutput == nil {
		panic("No target output or output layer to calculate error")
	}

	if d.output == nil {
		substracted := d.targetOutput.Subtract(d.value)

		return substracted
	} else {
		nextLayerErrors := d.output.Error()

		errors := make([]float64, d.size)

		for i := 0; i < d.size; i++ {
			errors[i] = 0.0

			for j := 0; j < d.output.Size(); j++ {
				errors[i] += d.output.Weights().At(i, j).Float() * nextLayerErrors.At(j).Float()
			}

			if errors[i] > 100 || errors[i] < -100 {
				log.Log("Error bigger than 100, %f, %f, %f", errors[i], nextLayerErrors.At(i).Float(), d.output.Weights().At(i, 0).Float())
				nextLayerErrors.PrintDebug()
			}
		}
		return tensor.NewTensor(tensor.FromArray(errors))
	}
}

func (d *Dense) Backward(learningRate float64) {
	/*
		for neuronIndex, neuron := range l.Neurons() {
			for _, input := range neuron.Inputs() {
				input.SetWeight(input.Weight() + learningRate*errors.At(neuronIndex).Float()*input.Input().Output()*
					activation.ActivationFunction(neuron.ActivationType()).Derivative(neuron.Output()))
			}
		}
	*/
	error := d.Error()

	if d.input != nil {
		for neuronIndex := 0; neuronIndex < d.size; neuronIndex++ {
			for neuronInputIndex := 0; neuronInputIndex < d.input.Size(); neuronInputIndex++ {
				index := tensor.NewIndex(neuronInputIndex, neuronIndex)

				delta := learningRate * error.At(neuronIndex).Float() *
					d.input.Output().At(0, neuronInputIndex).Float() *
					activation.ActivationFunction(d.activation).Derivative(
						d.value.At(0, neuronIndex).Float(),
					)

				log.Log("input output %f", d.input.Output().At(0, neuronInputIndex).Float())
				log.Log("error %f", error.At(neuronIndex).Float())
				log.Log("delta %f", delta)
				d.weights.Set(index, d.weights.Get(index).Float()+delta)
			}
		}

		d.input.Backward(learningRate)
	}
}

func (d *Dense) Nothing(f float64) {
	//
}

func (d *Dense) Set(inputValue *tensor.Tensor) {
	d.inputValue = inputValue
}
