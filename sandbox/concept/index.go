package concept

type Index interface {
	Get(Closure) (Variable, Interrupt)
	Set(Closure, Variable) Interrupt
	SubCodeBlockIterate(func(Index) bool) bool
	ToString
}
