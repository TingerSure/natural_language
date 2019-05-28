package lexer

type LexerInstanceGroup struct {
	instances []*LexerInstance
}

func (l *LexerInstanceGroup) AddInstance(instance *LexerInstance) {
	l.instances = append(l.instances, instance)
}

func (l *LexerInstanceGroup) GetInstances() []*LexerInstance {
	return l.instances
}

func NewLexerInstanceGroup() *LexerInstanceGroup {
	return &LexerInstanceGroup{}
}
