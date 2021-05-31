package concept

type Array interface {
	Variable
	Get(int) (Variable, Exception)
	Set(int, Variable) Exception
	Remove(int) Exception
	Append(Variable)
	Length() int
}
