package configuration

import "reflect"

type SingletonConfiguration[S any] struct {
	implementingType reflect.Type
	instance         any
}

func NewSingletonConfiguration[S any](implementingType reflect.Type, constructor func() S) SingletonConfiguration[S] {
	return SingletonConfiguration[S]{
		implementingType: implementingType,
	}
}

func (sc SingletonConfiguration[S]) GetOrCreateInstance() any {
	if sc.instance != nil {
		return sc.instance
	}

	panic("Not implemented.")
}
