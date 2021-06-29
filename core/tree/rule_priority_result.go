package tree

type PriorityResult struct {
	abandons []int
	result   int
}

func NewPriorityResult(result int) *PriorityResult {
	return &PriorityResult{
		result: result,
	}
}

func (a *PriorityResult) Result() int {
	return a.result
}

func (a *PriorityResult) Abandons() []int {
	return a.abandons
}

func (a *PriorityResult) AddAbandon(outsider int) *PriorityResult {
	for _, native := range a.abandons {
		if native == outsider {
			return a
		}
	}
	a.abandons = append(a.abandons, outsider)
	return a
}
