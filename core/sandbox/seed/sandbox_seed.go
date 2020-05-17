package seed

import (
	// "github.com/TingerSure/natural_language/core/sandbox/concept"
	// "github.com/TingerSure/natural_language/core/sandbox/variable"
	// "github.com/TingerSure/natural_language/core/sandbox"
)

type SandboxSeed struct {
	Variable *VariableSeed
}

func NewSandboxSeed() *SandboxSeed {
	return &SandboxSeed{
		Variable: NewVariableSeed(),
	}
}
