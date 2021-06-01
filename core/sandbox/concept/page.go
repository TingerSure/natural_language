package concept

type Page interface {
	Variable
	SetImport(String, Pipe) error
	SetPublic(String, Pipe) error
	SetPrivate(String, Pipe) error
}
