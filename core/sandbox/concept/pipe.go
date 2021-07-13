package concept

type Pipe interface {
	Type() string
	Get(Pool) (Variable, Interrupt)
	Set(Pool, Variable) Interrupt
	Call(Pool, Param) (Param, Exception)
	ToString
}
