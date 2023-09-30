package configuration

import "reflect"

type SingletonConfiguration struct {
	interfaceType    reflect.Type
	implementingType reflect.Type
}

func NewSingletonConfiguration(implementingType reflect.Type) SingletonConfiguration {
	return SingletonConfiguration{
		implementingType: implementingType,
	}
}

func (sc SingletonConfiguration) GetOrCreateInstance() any {
	panic("Not implemented.")
}
