package concept

type CodeBlock interface {
	Expression
	Size() int
	ToStringSimplify(string) string
	ToString(string) string
	AddStep(...Pipe)
	ExecWithInit(Pool, func(Pool) Interrupt) (Pool, Interrupt)
}
