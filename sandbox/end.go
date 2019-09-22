package sandbox

type End struct {
}

func (a *End) Exec(space *Closure) (bool, error) {
	return false, nil
}

func NewEnd() *End {
	return &End{}
}
