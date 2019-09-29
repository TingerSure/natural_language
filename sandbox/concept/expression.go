package concept

type Expression interface {
	Exec(Closure) Interrupt
	ToString(string) string
}
