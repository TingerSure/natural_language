package concept

type Variable interface {
	SetField(String, Variable) Exception
	GetField(String) (Variable, Exception)
	HasField(String) bool
	KeyField(String) String
	SizeField() int
	Iterate(func(String, Variable) bool) bool
	Call(String, Param) (Param, Exception)
	IsFunction() bool
	IsNull() bool
	Type() string
	ToString
}
