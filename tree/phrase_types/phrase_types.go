package phrase_types

/*
   The order here is very important, any type
   can't depend on the type behind it.
*/
const (
	Unknown string = "system.unknown"

	Target string = "system.target"

	AuxiliaryBelong string = "system.auxiliary.belong"
)
