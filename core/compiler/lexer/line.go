package lexer

import (
	"fmt"
	"strings"
)

type Line struct {
	content     []byte
	path        string
	rowStart    int
	colStart    int
	cursorStart int
	rowEnd      int
	colEnd      int
	cursorEnd   int
}

func NewLine(path string, content []byte) *Line {
	return &Line{
		content: content,
		path:    path,
	}
}

func NewLineFromTo(from, to *Line) *Line {
	if from == to {
		return from
	}
	return &Line{
		content:     from.content,
		path:        from.path,
		rowStart:    from.rowStart,
		colStart:    from.colStart,
		cursorStart: from.cursorStart,
		rowEnd:      to.rowEnd,
		colEnd:      to.colEnd,
		cursorEnd:   to.cursorEnd,
	}
}

func (l *Line) Path() string {
	return l.path
}

func (l *Line) Content() []byte {
	return l.content
}

func (l *Line) GetStart() (int, int, int) {
	return l.rowStart, l.colStart, l.cursorStart
}

func (l *Line) GetEnd() (int, int, int) {
	return l.rowEnd, l.colEnd, l.cursorEnd
}

func (l *Line) SetStart(row, col, cursor int) {
	l.rowStart = row
	l.colStart = col
	l.cursorStart = cursor
}

func (l *Line) SetEnd(row, col, cursor int) {
	l.rowEnd = row
	l.colEnd = col
	l.cursorEnd = cursor
}

func (l *Line) End() *Line {
	line := NewLine(l.path, l.content)
	line.SetStart(l.rowEnd, l.colEnd, l.cursorEnd)
	line.SetEnd(l.rowEnd, l.colEnd, l.cursorEnd)
	return line
}

func (l *Line) ToPath() string {
	return fmt.Sprintf("%v:%v:%v", l.path, l.rowStart+1, l.colStart+1)
}

func (l *Line) ToLine() string {
	prev := strings.LastIndex(string(l.content[:l.cursorStart]), "\n") + 1
	next := strings.Index(string(l.content[l.cursorStart:]), "\n") + l.cursorStart
	tabSize := strings.Count(string(l.content[prev:l.cursorStart]), "\t")
	return fmt.Sprintf(
		"%v\n%v%v^",
		string(l.content[prev:next]),
		strings.Repeat("\t", tabSize),
		strings.Repeat(" ", l.cursorStart-prev-tabSize),
	)
}

func (l *Line) ToString() string {
	return fmt.Sprintf("%v:\n%v", l.ToPath(), l.ToLine())
}
