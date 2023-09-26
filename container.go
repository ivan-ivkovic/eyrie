package falconsnest

import (
	"falconsnest/internal/configuration"
	"reflect"
)

type Container interface {
	addConfiguration(t reflect.Type, config configuration.Configuration)
}

type container struct {
}

var c container

func GetContainer() Container {
	return c
}

func (c container) addConfiguration(t reflect.Type, config configuration.Configuration) {

}

func Register[I any, S any]() Registrar[I, S] {
	return newRegistrar[I, S](c)
}
