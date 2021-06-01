package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/creator"
)

type LibraryManager struct {
	pages   map[string]concept.Pipe
	Sandbox *creator.SandboxCreator
	sources []Source
}

func (p *LibraryManager) GetSources() []Source {
	return p.sources
}

func (p *LibraryManager) AddSource(source Source) {
	p.sources = append(p.sources, source)
}

func (l *LibraryManager) GetPage(name string) concept.Pipe {
	return l.pages[name]
}

func (l *LibraryManager) AddPage(name string, page concept.Pipe) {
	l.pages[name] = page
}

func NewLibraryManager() *LibraryManager {
	return &LibraryManager{
		pages:   map[string]concept.Pipe{},
		Sandbox: creator.NewSandboxCreator(),
		sources: []Source{},
	}
}
