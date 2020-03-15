package concept

type Closure interface {
	SetParent(Closure)
	SetReturn(String, Variable)
	MergeReturn(Closure)
	IterateReturn(func(String, Variable) bool) bool
	InitLocal(String, Variable)
	GetLocal(String) (Variable, Interrupt)
	SetLocal(String, Variable) Interrupt
	GetBubble(String) (Variable, Interrupt)
	SetBubble(String, Variable) Interrupt
	AddExtempore(Index, Variable)
	IterateHistory(func(String, Variable) bool) bool
	IterateExtempore(func(Index, Variable) bool) bool
	IterateLocal(func(String, Variable) bool) bool
	IterateBubble(func(String, Variable) bool) bool
	Clear()
}
