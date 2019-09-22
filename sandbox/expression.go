package sandbox

type Expression interface {
	Exec(*Closure) (bool, error)
}
