package sandbox

type Return struct {
	key    string
	result Index
}

func (a *Return) Exec(space *Closure) (bool, error) {
	result, err := a.result.Get(space)

	if err != nil {
		return false, err
	}
	space.SetReturn(a.key, result)
	return true, nil
}

func NewReturn(key string, result Index) *Return {
	return &Return{
		key:    key,
		result: result,
	}
}
