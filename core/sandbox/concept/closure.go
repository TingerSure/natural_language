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
	AddExtempore(Index, Variable)
	IterateHistory(func(string, Variable) bool) bool
	IterateExtempore(func(Index, Variable) bool) bool
	IterateLocal(func(string, Variable) bool) bool
	IterateBubble(func(string, Variable) bool) bool
	Clear()
}
