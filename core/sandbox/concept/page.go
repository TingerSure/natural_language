package concept

type Page interface {
	Variable
	SetImport(String, Pipe) Exception
	SetPublic(String, Pipe) Exception
	SetPrivate(String, Pipe) Exception
}
