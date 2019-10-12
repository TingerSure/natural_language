package concept

type Param interface {
	Variable
	Set(key string, value Variable)
	Get(key string) Variable
	Init(params map[string]Variable)
}
