package tree

type StructRuleParam struct {
	Size   int
	Create func() Phrase
	From   string
	Types  []*PhraseType
}

type StructRule struct {
	param *StructRuleParam
}

func (r *StructRule) Size() int {
	return r.param.Size
}

func (r *StructRule) Types() []*PhraseType {
	return r.param.Types
}

func (r *StructRule) GetFrom() string {
	return r.param.From
}

func (r *StructRule) Match(treasures []Phrase) bool {
	if len(treasures) < r.param.Size {
		return false
	}
	match := true
	for index, treasure := range treasures[len(treasures)-r.param.Size:] {
		match = match && r.param.Types[index].Match(treasure.Types())
	}
	return match
}

func (r *StructRule) Create(treasures []Phrase) Phrase {
	new := r.param.Create()
	for index, treasure := range treasures[len(treasures)-r.param.Size:] {
		new.SetChild(index, treasure)
	}
	return new
}

func NewStructRule(param *StructRuleParam) *StructRule {
	param.Size = len(param.Types)
	if param.Size == 0 {
		panic("There must be at least one type here.")
	}
	if param.Create == nil {
		panic("no create function in this struct rule!")
	}
	return &StructRule{
		param: param,
	}
}
