package concept

type Page interface {
	Variable
	SetImport(String, Index) error
	SetPublic(String, Index) error
	SetPrivate(String, Index) error
}
