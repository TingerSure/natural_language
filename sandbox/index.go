package sandbox

type Index interface {
	Get(*Closure) (Variable, *Exception)
	Set(*Closure, Variable) *Exception
	ToString(prefix string) string
}
