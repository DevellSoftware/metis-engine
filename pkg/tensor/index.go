package tensor

import "fmt"

type Index struct {
	index []int
}

func NewIndex(index ...int) Index {
	return Index{index: index}
}

func (i Index) Hash() string {
	str := fmt.Sprintf("%v", i.index)

	return str
}

func (i Index) Parts() []int {
	return i.index
}

func (i Index) Equals(other Index) bool {
	normalizedIndex1 := i.Normalize()
	normalizedIndex2 := other.Normalize()

	if len(normalizedIndex1.index) != len(normalizedIndex2.index) {
		return false
	}

	for i, part := range normalizedIndex1.index {
		if part != normalizedIndex2.index[i] {
			return false
		}
	}

	return true
}

func (i Index) String() string {
	return fmt.Sprintf("<%v>", i.index)
}

func (i Index) PadToRank(rank int) Index {
	if len(i.index) == rank {
		return i
	}

	padded := make([]int, rank)
	for i, part := range i.index {
		padded[i] = part
	}

	return NewIndex(padded...)
}

func (i Index) Normalize() Index {
	return i.PadToRank(4)
}
