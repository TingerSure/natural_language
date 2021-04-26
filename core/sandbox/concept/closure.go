package concept

type Closure interface {
	SetParent(Closure)
	SetReturn(String, Variable)
	MergeReturn(Closure)
	IterateReturn(func(String, Variable) bool) bool
	InitLocal(String, Variable)
	PeekLocal(String) (Variable, Exception)
	GetLocal(String) (Variable, Exception)
	HasLocal(String) bool
	SetLocal(String, Variable) Exception
	PeekBubble(String) (Variable, Exception)
	GetBubble(String) (Variable, Exception)
	HasBubble(String) bool
	SetBubble(String, Variable) Exception
	AddExtempore(Index, Variable)
	IterateHistory(func(String, Variable) bool) bool
	IterateExtempore(func(Index, Variable) bool) bool
	IterateLocal(func(String, Variable) bool) bool
	IterateBubble(func(String, Variable) bool) bool
	Clear()
}
