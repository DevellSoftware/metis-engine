package activation

type ActivationType int

const (
	SigmoidFunction ActivationType = iota
	ReLUFunction
	LinearFunction
)

func ActivationFunction(t ActivationType) Activation {
	switch t {
	case SigmoidFunction:
		return Sigmoid{}
	case ReLUFunction:
		return ReLU{}
	case LinearFunction:
		return Linear{}
	default:
		return Sigmoid{}
	}
}
