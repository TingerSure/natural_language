package sandbox

type Expression interface {
	Exec(*Closure) Interrupt
	ToString(prefix string) string
}
