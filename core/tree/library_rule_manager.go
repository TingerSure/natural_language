package tree

type VocabularyRuleManager interface {
	AddRule(rules *VocabularyRule)
	RemoveRule(need func(rule *VocabularyRule) bool)
}

type StructRuleManager interface {
	AddRule(rules *StructRule)
	RemoveRule(need func(rule *StructRule) bool)
}

type PriorityRuleManager interface {
	AddRule(rules *PriorityRule)
	RemoveRule(need func(rule *PriorityRule) bool)
}

type TypesManager interface {
	AddTypes(values *PhraseType)
	RemoveTypes(need func(types *PhraseType) bool)
}

type DutyRuleManager interface {
	AddRule(rules *DutyRule)
	RemoveRule(need func(rule *DutyRule) bool)
}
