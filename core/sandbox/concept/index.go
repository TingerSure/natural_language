package concept

type Index interface {
	Type() string
	Anticipate(Closure) Variable
	Get(Closure) (Variable, Interrupt)
	Set(Closure, Variable) Interrupt
	SubCodeBlockIterate(func(Index) bool) bool
	ToString
}
