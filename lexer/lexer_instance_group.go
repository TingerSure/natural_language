package lexer

type LexerInstanceGroup struct {
	instances []*LexerInstance
}

func (l *LexerInstanceGroup) Size() int {
	return len(l.instances)
}

func (l *LexerInstanceGroup) AddInstance(instance *LexerInstance) {
	l.instances = append(l.instances, instance)
}

func (l *LexerInstanceGroup) GetInstances() []*LexerInstance {
	return l.instances
}

func (l *LexerInstanceGroup) FilterNull() *LexerInstanceGroup {
	var noNullGroup *LexerInstanceGroup = NewLexerInstanceGroup()
	for _, instance := range l.instances {
		if instance.HasNull() {
			continue
		}
		noNullGroup.AddInstance(instance)
	}
	return noNullGroup
}

func NewLexerInstanceGroup() *LexerInstanceGroup {
	return &LexerInstanceGroup{}
}
