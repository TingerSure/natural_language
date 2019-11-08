package phrase_types

const (
	Unknown string = "system.unknown"
	Number  string = "system.number"

	/*
		Operator_Result = Operator_Left Operator Operator_Right
	*/
	Operator        string = "system.operator"
	Operator_Left   string = "left"
	Operator_Right  string = "right"
	Operator_Result string = "result"

	Event           string = "system.event"
	Target          string = "system.target"
	AuxiliaryBelong string = "system.auxiliary.belong"
	Action          string = "system.action"
)
