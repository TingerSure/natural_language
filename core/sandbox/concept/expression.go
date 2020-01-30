package concept

type Expression interface {
	Index
	Exec(Closure) (Variable, Interrupt)
}
