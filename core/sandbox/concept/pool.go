package concept

type Pool interface {
	Variable
	SetParent(Pool)
	InitLocal(String, Variable)
	KeyLocal(String) String
	PeekLocal(String) (Variable, Exception)
	GetLocal(String) (Variable, Exception)
	HasLocal(String) bool
	SetLocal(String, Variable) Exception
	IterateLocal(func(String, Variable) bool) bool
	KeyBubble(String) String
	PeekBubble(String) (Variable, Exception)
	GetBubble(String) (Variable, Exception)
	HasBubble(String) bool
	SetBubble(String, Variable) Exception
	IterateBubble(func(String, Variable) bool) bool
	AddExtempore(Function, Variable)
	IterateHistory(func(String, Variable) bool) bool
	IterateExtempore(func(Function, Variable) bool) bool
	Clear()
}
