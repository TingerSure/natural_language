package concept

type Param interface {
	Variable
	Set(String, Variable) Param
	Get(String) Variable
	Init(func(func(String, Variable) bool) bool)
}
