package sandbox

type Expression interface {
	Exec(*Closure) Interrupt
}
