package tree

type Library interface {
	GetPage(key string) *Page
	SetPage(key string, value *Page) Library
}
