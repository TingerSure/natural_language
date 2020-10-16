package tree

type Abandon struct {
	Offset int
	Value  Phrase
}

type AbandonGroup struct {
	values []*Abandon
}

func NewAbandonGroup() *AbandonGroup {
	return &AbandonGroup{}
}

func (a *AbandonGroup) Values() []*Abandon {
	return a.values
}

func (a *AbandonGroup) Add(comer *Abandon) *AbandonGroup {
	for _, value := range a.values {
		if value.Offset == comer.Offset && value.Value == comer.Value {
			return a
		}
	}
	a.values = append(a.values, comer)
	return a
}

func (a *AbandonGroup) Merge(group *AbandonGroup) {
	if group == nil {
		return
	}

	for index := 0; index < len(group.values); index++ {
		a.Add(group.values[index])
	}
}
