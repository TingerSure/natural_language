package concept

type Index interface {
	Get(Closure) (Variable, Interrupt)
	Set(Closure, Variable) Interrupt
	ToString(string) string
}
