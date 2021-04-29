package concept

type Page interface {
	Variable
	SetImport(String, Index) error
	SetExport(String, Index) error
	SetVar(String, Index) error
}
