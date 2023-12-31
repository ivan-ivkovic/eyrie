package configuration

import (
	"reflect"
)

type SingletonConfiguration[T any] struct {
	implementingType reflect.Type
	constructor      func() T
	instance         *T
	resolving        bool
}

func NewSingletonConfiguration[T any](implementingType reflect.Type, constructor func() T) *SingletonConfiguration[T] {
	return &SingletonConfiguration[T]{
		implementingType: implementingType,
		constructor:      constructor,
	}
}

func (sc *SingletonConfiguration[T]) GetOrCreateInstance() (any, error) {
	if sc.resolving {
		return nil, newConfigurationError(RecursiveConstruction)
	}
	sc.resolving = true
	defer func() { sc.resolving = false }()

	if sc.instance == nil {
		var instance T = sc.constructor()
		sc.instance = &instance
	}

	return any(*sc.instance), nil
}
