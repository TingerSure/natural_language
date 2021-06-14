package message

import (
	"fmt"
)

type Message struct {
	path string
	line string
}

func NewMessage(path string, line string) *Message {
	return &Message{
		path,
		line,
	}
}

func (l *Message) ToPath() string {
	return l.path
}

func (l *Message) ToLine() string {
	return l.line
}

func (l *Message) ToString() string {
	return fmt.Sprintf("%v:\n%v", l.path, l.line)
}
