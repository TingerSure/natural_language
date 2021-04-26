package concept

type Page interface {
	Variable
	SetImport(String, Index) Exception
	SetExport(String, Index) Exception
	SetVar(String, Index) Exception
}
