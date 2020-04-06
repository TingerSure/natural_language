package tree

type Library interface {
	PageIterate(func(Page) bool) bool
	GetPage(key string) Page
	SetPage(key string, value Page) Library
}
