package configuration

import "reflect"

type TransientConfiguration struct {
	interfaceType    reflect.Type
	implementingType reflect.Type
}

func NewTransientConfiguration(implementingType reflect.Type) TransientConfiguration {
	return TransientConfiguration{
		implementingType: implementingType,
	}
}

func (tc TransientConfiguration) GetOrCreateInstance() any {
	panic("Not implemented.")
}
