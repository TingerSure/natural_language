package runtime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Scan struct {
	listener func(string)
	prompt   func()
	reader   *os.File
}

func (s *Scan) Run() {
	bufferReader := bufio.NewReader(s.reader)

	for {
		s.prompt()
		input, err := bufferReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			} else {
				fmt.Println("error -> Read file error!", err)
				return
			}
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		s.listener(input)
	}
}

var (
	defaultPrompt = func() {
		// Do nothing
	}
)

func NewScan(reader *os.File, listener func(string), prompt func()) *Scan {
	if prompt == nil {
		prompt = defaultPrompt
	}
	return &Scan{
		reader:   reader,
		listener: listener,
		prompt:   prompt,
	}
}
