package tensor

import "fmt"

type Value interface {
	Float() float64
	Length() int
	String() string
}

type NumberValue struct {
	value float64
}

func NewNumberValue(value float64) Value {
	return &NumberValue{value: value}
}

func (n *NumberValue) Float() float64 {
	return n.value
}

func (n *NumberValue) String() string {
	return fmt.Sprintf("%f", n.value)
}

func NewValue(value interface{}) Value {
	switch value := value.(type) {
	case float64:
		return NewNumberValue(value)
	default:
		panic("unhandled type")
	}
}

func (v *NumberValue) Length() int {
	return 1
}

type TensorValue struct {
	values Tensor
}

func NewTensorValue(values Tensor) Value {
	return &TensorValue{values: values}
}

func (t *TensorValue) Float() float64 {
	panic("not implemented")
}

func (t *TensorValue) Length() int {
	return len(t.values.values)
}

func (t *TensorValue) String() string {
	return fmt.Sprintf("[%s]", t.values.ValuesString())
}

type UndefinedValue struct{}

func NewUndefinedValue() Value {
	return &UndefinedValue{}
}

func (u *UndefinedValue) Float() float64 {
	return 0
}

func (u *UndefinedValue) Length() int {
	return 0
}

func (u *UndefinedValue) String() string {
	return "undefined"
}
