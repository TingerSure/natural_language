package grammar

import (
	"fmt"
	"strings"
)

type TableClosure struct {
	projects map[*TableProject]map[Symbol]bool
	id       int
}

func NewTableClosure(id int) *TableClosure {
	return &TableClosure{
		projects: map[*TableProject]map[Symbol]bool{},
		id:       id,
	}
}

func (t *TableClosure) NextChildren() map[Symbol]bool {
	children := map[Symbol]bool{}
	for project, _ := range t.projects {
		if project.IsEnd() {
			continue
		}
		children[project.GetNextChild()] = true
	}
	return children
}

func (t *TableClosure) Include(another *TableClosure) bool {
	if len(t.projects) != len(another.projects) {
		return false
	}
	for project, lookaheads := range t.projects {
		anotherLookaheads := another.projects[project]
		if anotherLookaheads == nil || len(lookaheads) < len(anotherLookaheads) {
			return false
		}
		for symbol, _ := range anotherLookaheads {
			if !lookaheads[symbol] {
				return false
			}
		}
	}
	return true
}

func (t *TableClosure) AddProject(project *TableProject, lookaheads map[Symbol]bool) (success map[Symbol]bool) {
	if t.projects[project] == nil {
		t.projects[project] = map[Symbol]bool{}
	}
	success = map[Symbol]bool{}
	for symbol, _ := range lookaheads {
		if !t.projects[project][symbol] {
			t.projects[project][symbol] = true
			success[symbol] = true
		}
	}
	return
}

func (t *TableClosure) Id() int {
	return t.id
}

func (t *TableClosure) Size() int {
	return len(t.projects)
}

func (t *TableClosure) GetProjects() map[*TableProject]map[Symbol]bool {
	return t.projects
}

func (t *TableClosure) GetProjectsByNextChild(nextChild Symbol) map[*TableProject]map[Symbol]bool {
	projects := map[*TableProject]map[Symbol]bool{}
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
		for symbol, _ := range lookaheads {
			symbolNames = append(symbolNames, symbol.Name())
		}
		values = append(values, fmt.Sprintf("|%v|%v|", project.ToString(), strings.Join(symbolNames, " , ")))
	}
	return strings.Join(values, "\n")
}
