package concept

type Param interface {
	Variable
	Set(String, Variable)
	Get(String) Variable
	// SetIndex(int, Variable)
	// GetIndex(int) Variable
	Copy() Param
	Iterate(func(String, Variable) bool) bool
}
