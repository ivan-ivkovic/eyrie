package tests

type IInterface interface{}

type Struct struct {
	memoryAddress *int
}

func NewStruct() Struct {
	return Struct{
		memoryAddress: new(int),
	}
}
