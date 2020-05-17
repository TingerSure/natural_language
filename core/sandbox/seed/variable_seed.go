package seed

import (
	// "github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type VariableSeed struct {
	String *variable.StringSeed
}

func NewVariableSeed() *VariableSeed {
	return &VariableSeed{
		String: variable.NewStringSeed(),
	}
}
