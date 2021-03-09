package grammar

type TableClosure struct {
	projects map[*TableProject]bool
	id       int
}

func NewTableClosure(id int) *TableClosure {
	return &TableClosure{
		projects: map[*TableProject]bool{},
		id:       id,
	}
}

func (t *TableClosure) AddProject(project *TableProject) bool {
	if t.projects[project] {
		return false
	}
	t.projects[project] = true
	return true
}

func (t *TableClosure) Id() int {
	return t.id
}

func (t *TableClosure) Size() int {
	return len(t.projects)
}

func (t *TableClosure) GetProjects() map[*TableProject]bool {
	return t.projects
}
