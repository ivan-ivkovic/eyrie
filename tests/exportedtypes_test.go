package tests

type IInterface interface {
	Method1()
}

type Struct struct {
	memoryAddress *int
}

func NewStruct() Struct {
	return Struct{
		memoryAddress: new(int),
	}
}

func (s Struct) Method1() {
	panic("Not implemented.")
}
