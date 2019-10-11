package concept

type Expression interface {
	Exec(Closure) (Variable, Interrupt)
	ToString(string) string
}
