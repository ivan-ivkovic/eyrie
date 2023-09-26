package falconsnest

import (
	"falconsnest/internal/configuration"
	"fmt"
	"reflect"
)

type Registrar[I any, S any] interface {
	AsSingleton()
	AsTransient()
}

type registrar[I any, S any] struct {
	container Container
}

func newRegistrar[I any, S any](c Container) Registrar[I, S] {
	return registrar[I, S]{
		container: c,
	}
}

func (r registrar[I, S]) AsSingleton() {
	interfaceType, implementingType := r.validateTypes[I, S]()
	config := configuration.NewSingletonConfiguration(interfaceType, implementingType)

	r.container.addConfiguration(interfaceType, config)
}

func (r registrar[I, S]) AsTransient() {
	interfaceType, implementingType := r.validateTypes[I, S]()
	config := configuration.NewTransientConfiguration(interfaceType, implementingType)

	r.container.addConfiguration(interfaceType, config)
}

func (r registrar[I, S]) validateTypes() (reflect.Type, reflect.Type) {
	var i [0]I
	var s [0]S
	var it = reflect.TypeOf(i).Elem()
	var st = reflect.TypeOf(s).Elem()
	var ik = it.Kind()
	var sk = st.Kind()

	if ik != reflect.Interface || sk != reflect.Struct {
		panic(newRegistrationError("Interface and struct expected as type parameters."))
	}

	if !st.Implements(it) {
		var msg = fmt.Sprintf("%s does not implement %s.", st.Name(), it.Name())
		panic(newRegistrationError(msg))
	}

	return it, st
}
