package concept

type Closure interface {
	SetParent(Closure)
	SetReturn(string, Variable)
	MergeReturn(Closure)
	Return() map[string]Variable
	InitLocal(string)
	GetLocal(string) (Variable, Interrupt)
	SetLocal(string, Variable) Interrupt
	GetBubble(string) (Variable, Interrupt)
	SetBubble(string, Variable) Interrupt
	GetCache(int) Variable
	SetCache(int, Variable)
	Clear()
}
