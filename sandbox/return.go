package sandbox

type Return struct {
	key    string
	result Index
}

func (a *Return) Exec(space *Closure) Interrupt {
	result, exception := a.result.Get(space)

	if exception != nil {
		return exception
	}
	space.SetReturn(a.key, result)
	return nil
}

func NewReturn(key string, result Index) *Return {
	return &Return{
		key:    key,
		result: result,
	}
}
