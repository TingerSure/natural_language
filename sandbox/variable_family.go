package sandbox

type VariableFamily struct {
}

func (v *VariableFamily) IsNumber(value Variable) (*Number, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableNumberType {
		number, yes := value.(*Number)
		return number, yes
	}
	return nil, false
}

func newVariableFamily() *VariableFamily {
	return &VariableFamily{}
}

var (
	VariableFamilyInstance *VariableFamily = newVariableFamily()
)
