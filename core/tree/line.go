package tree

import (
	"fmt"
)

type Line struct {
	path string
	line string
}

func NewLine(path string, line string) *Line {
	return &Line{
		path,
		line,
	}
}

func (l *Line) ToPath() string {
	return l.path
}

func (l *Line) ToLine() string {
	return l.line
}

func (l *Line) ToString() string {
	return fmt.Sprintf("%v:\n%v", l.path, l.line)
}
