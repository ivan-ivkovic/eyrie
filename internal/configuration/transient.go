package configuration

import "reflect"

type TransientConfiguration[T any] struct {
	implementingType reflect.Type
}

func NewTransientConfiguration[T any](implementingType reflect.Type, constructor func() T) TransientConfiguration[T] {
	return TransientConfiguration[T]{
		implementingType: implementingType,
	}
}

func (tc TransientConfiguration[T]) GetOrCreateInstance() any {
	panic("Not implemented.")
}
