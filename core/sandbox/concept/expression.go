package concept

type Expression interface {
	Pipe
	Exec(Pool) (Variable, Interrupt)
}
