package runtime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Scan struct {
	onReader     func(string) bool
	beforeReader func()
	stream       *os.File
}

func (s *Scan) Run() {
	bufferReader := bufio.NewReader(s.stream)

	for {
		s.beforeReader()
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
		if !s.onReader(input) {
			break
		}
	}
}

type ScanParam struct {
	OnReader     func(string) bool
	BeforeReader func()
	Stream       *os.File
}

func NewScan(param *ScanParam) *Scan {
	return &Scan{
		stream:       param.Stream,
		onReader:     param.OnReader,
		beforeReader: param.BeforeReader,
	}
}
