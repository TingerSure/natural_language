package concept

type CodeBlock interface {
	Expression
	Size() int
	ToStringSimplify(string) string
	ToString(string) string
	Steps() []Pipe
	AddStep(...Pipe)
	ExecWithInit(Pool, func(Pool) Interrupt) (Pool, Interrupt)
}
