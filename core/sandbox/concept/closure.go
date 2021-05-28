package concept

type Closure interface {
	SetParent(Closure)
	InitLocal(String, Variable)
	KeyLocal(String) String
	PeekLocal(String) (Variable, Exception)
	GetLocal(String) (Variable, Exception)
	HasLocal(String) bool
	SetLocal(String, Variable) Exception
	KeyBubble(String) String
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
