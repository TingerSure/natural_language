package concept

type Param interface {
	Variable
	Set(String, Variable) Param
	Get(String) Variable
	Copy() Param
	Init(func(func(String, Variable) bool) bool) Param
	Iterate(func(String, Variable) bool) bool
}
