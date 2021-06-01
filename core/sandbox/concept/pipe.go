package concept

type Pipe interface {
	Type() string
	Anticipate(Pool) Variable
	Get(Pool) (Variable, Interrupt)
	Set(Pool, Variable) Interrupt
	Call(Pool, Param) (Param, Exception)
	CallAnticipate(Pool, Param) Param
	ToString
}
