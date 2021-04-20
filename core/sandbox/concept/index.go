package concept

type Index interface {
	Type() string
	Anticipate(Closure) Variable
	Get(Closure) (Variable, Interrupt)
	Set(Closure, Variable) Interrupt
	Call(Closure, Param) (Param, Exception)
	CallAnticipate(Closure, Param) Param
	SubCodeBlockIterate(func(Index) bool) bool
	ToString
}
