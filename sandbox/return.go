package sandbox

import (
	"fmt"
)

type Return struct {
	key    string
	result Index
}

func (a *Return) ToString(prefix string) string {
	return fmt.Sprintf("%vreturn[%v] %v", prefix, a.key, a.result.ToString(prefix))
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
