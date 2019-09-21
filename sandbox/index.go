package sandbox

type Index interface {
	Get(*Closure) (Variable, error)
	Set(*Closure, Variable) error
}
