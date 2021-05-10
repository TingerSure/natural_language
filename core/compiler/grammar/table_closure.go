package grammar

import (
	"fmt"
	"strings"
)

type TableClosure struct {
	projects map[*TableProject]*SymbolSet
	id       int
}

func NewTableClosure(id int) *TableClosure {
	return &TableClosure{
		projects: map[*TableProject]*SymbolSet{},
		id:       id,
	}
}

func (t *TableClosure) NextChildren() *SymbolSet {
	children := NewSymbolSet()
	for project, _ := range t.projects {
		if project.IsEnd() {
			continue
		}
		children.Add(project.GetNextChild())
	}
	return children
}

func (t *TableClosure) Include(another *TableClosure) bool {
	if len(t.projects) != len(another.projects) {
		return false
	}
	for project, lookaheads := range t.projects {
		anotherLookaheads := another.projects[project]
		if anotherLookaheads == nil || lookaheads.Size() < anotherLookaheads.Size() {
			return false
		}
		if anotherLookaheads.Iterate(func(symbol Symbol) bool {
			return !lookaheads.Has(symbol)
		}) {
			return false
		}
	}
	return true
}

func (t *TableClosure) AddProject(project *TableProject, lookaheads *SymbolSet) *SymbolSet {
	if t.projects[project] == nil {
		t.projects[project] = NewSymbolSet()
	}
	success := NewSymbolSet()

	lookaheads.Iterate(func(symbol Symbol) bool {
		if !t.projects[project].Has(symbol) {
			t.projects[project].Add(symbol)
			success.Add(symbol)
		}
		return false
	})
	return success
}

func (t *TableClosure) Id() int {
	return t.id
}

func (t *TableClosure) Size() int {
	return len(t.projects)
}

func (t *TableClosure) GetProjects() map[*TableProject]*SymbolSet {
	return t.projects
}

func (t *TableClosure) GetProjectsByNextChild(nextChild Symbol) map[*TableProject]*SymbolSet {
	projects := map[*TableProject]*SymbolSet{}
	for project, lookaheads := range t.projects {
		if project.GetNextChild() == nextChild {
			projects[project] = lookaheads
		}
	}
	return projects
}

func (t *TableClosure) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("|%v|lookaheads| ", t.id))
	values = append(values, "|:--:|:--:|")
	for project, lookaheads := range t.projects {
		symbolNames := []string{}
		lookaheads.Iterate(func(symbol Symbol) bool {
			symbolNames = append(symbolNames, symbol.Name())
			return false
		})
		values = append(values, fmt.Sprintf("|%v|%v|", project.ToString(), strings.Join(symbolNames, " , ")))
	}
	return strings.Join(values, "\n")
}
