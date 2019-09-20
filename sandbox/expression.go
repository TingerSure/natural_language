package sandbox

type Expression interface {
	Exec(Closure) error
}
