package grammar

type RuleSet struct {
	values map[Symbol]map[*Rule]bool
	size   int
}

func NewRuleSet() *RuleSet {
	return &RuleSet{
		values: map[Symbol]map[*Rule]bool{},
		size:   0,
	}
}

func (s *RuleSet) Iterate(on func(*Rule) bool) bool {
	for _, rules := range s.values {
		for rule, _ := range rules {
			if on(rule) {
				return true
			}
		}
	}
	return false
}

func (s *RuleSet) Size() int {
	return s.size
}

func (s *RuleSet) Add(rules ...*Rule) {
	for _, rule := range rules {
		s.addStep(rule)
	}
}

func (s *RuleSet) addStep(rule *Rule) {
	rules := s.values[rule.GetResult()]
	if rules == nil {
		rules = map[*Rule]bool{}
		s.values[rule.GetResult()] = rules
	}
	if rules[rule] {
		return
	}
	rules[rule] = true
	s.size++
}

func (s *RuleSet) Remove(rule *Rule) {
	rules := s.values[rule.GetResult()]
	if rules == nil {
		return
	}
	if !rules[rule] {
		return
	}
	delete(rules, rule)
	s.size--
}

func (s *RuleSet) HasByResultEmpty(symbol Symbol) bool {
	rules := s.values[symbol]
	if rules == nil {
		return false
	}
	for rule := range rules {
		if rule.Size() == 0 {
			return true
		}
	}
	return false
}

func (s *RuleSet) HasByResult(symbol Symbol) bool {
	rules := s.values[symbol]
	if rules == nil {
		return false
	}
	return len(rules) > 0
}

func (s *RuleSet) Has(rule *Rule) bool {
	rules := s.values[rule.GetResult()]
	if rules == nil {
		return false
	}
	return rules[rule]
}
