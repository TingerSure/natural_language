package lexer

type FlowGroup struct {
	instances []*Flow
}

func (l *FlowGroup) Size() int {
	return len(l.instances)
}

func (l *FlowGroup) AddInstance(instance *Flow) {
	l.instances = append(l.instances, instance)
}

func (l *FlowGroup) GetInstances() []*Flow {
	return l.instances
}

func (l *FlowGroup) FilterNull() *FlowGroup {
	var noNullGroup *FlowGroup = NewFlowGroup()
	for _, instance := range l.instances {
		if instance.HasNull() {
			continue
		}
		noNullGroup.AddInstance(instance)
	}
	return noNullGroup
}

func NewFlowGroup() *FlowGroup {
	return &FlowGroup{}
}
