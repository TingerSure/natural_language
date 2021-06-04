package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/creator"
)

type LibraryManager struct {
	pages       map[string]concept.Pipe
	Sandbox     *creator.SandboxCreator
	Vocabularys VocabularyRuleManager
	Structs     StructRuleManager
	Priorities  PriorityRuleManager
	Types       TypesManager
	Duties      DutyRuleManager
}

func (l *LibraryManager) GetPage(name string) concept.Pipe {
	return l.pages[name]
}

func (l *LibraryManager) AddPage(name string, page concept.Pipe) {
	l.pages[name] = page
}

func NewLibraryManager(
	Sandbox *creator.SandboxCreator,
	Vocabularys VocabularyRuleManager,
	Structs StructRuleManager,
	Priorities PriorityRuleManager,
	Types TypesManager,
	Duties DutyRuleManager,
) *LibraryManager {
	return &LibraryManager{
		pages:       map[string]concept.Pipe{},
		Sandbox:     Sandbox,
		Vocabularys: Vocabularys,
		Structs:     Structs,
		Priorities:  Priorities,
		Types:       Types,
		Duties:      Duties,
	}
}
