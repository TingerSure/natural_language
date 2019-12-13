package concept

type Index interface {
	Get(Closure) (Variable, Interrupt)
	Set(Closure, Variable) Interrupt
	SubIterate(func(Index) bool) bool
	ToString
}
