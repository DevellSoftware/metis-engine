package tensor

type TensorConstruct func(t *Tensor)

func FromArray(array []float64) TensorConstruct {
	return func(t *Tensor) {
		t.rank = 1

		for i, value := range array {
			t.Set(NewIndex(i), value)
		}
	}
}

func FromArrayRank1(array []float64) TensorConstruct {
	return FromArray(array)
}

func FromArrayRank2(array [][]float64) TensorConstruct {
	return func(t *Tensor) {
		t.rank = 2

		for i, row := range array {
			for j, value := range row {
				t.Set(NewIndex(i, j), value)
			}
		}
	}
}

func WithRank(rank int) TensorConstruct {
	return func(t *Tensor) {
		t.rank = rank
	}
}
