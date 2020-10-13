package runtime

import (
	"errors"
	"github.com/TingerSure/natural_language/core/parser"
	"github.com/TingerSure/natural_language/core/sandbox"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
)

var (
	runtimeStructErrorFormatDefault = func(road *parser.Road) string {
		return "No struct rule can match this sentence."
	}

	runtimePriorityErrorFormatDefault = func(road *parser.Road) string {

		return "No priority rule can distinguish all meanings"
	}
)

type Runtime struct {
	parser              *parser.Parser
	libs                *LibraryManager
	box                 *sandbox.Sandbox
	rootSpace           *closure.Closure
	defaultLanguage     string
	structErrorFormat   func(*parser.Road) string
	priorityErrorFormat func(*parser.Road) string
}

func (r *Runtime) SetStructErrorFormat(format func(*parser.Road) string) {
	if format == nil {
		r.structErrorFormat = runtimeStructErrorFormatDefault
	}
	r.structErrorFormat = format
}

func (r *Runtime) SetPriorityErrorFormat(format func(*parser.Road) string) {
	if format == nil {
		r.priorityErrorFormat = runtimePriorityErrorFormatDefault
	}
	r.priorityErrorFormat = format
}

func (r *Runtime) SetDefaultLanguage(name string) {
	r.defaultLanguage = name
}

func (r *Runtime) GetDefaultLanguage() string {
	return r.defaultLanguage
}

func (r *Runtime) GetLibraryManager() *LibraryManager {
	return r.libs
}

func (r *Runtime) Bind() {
	r.libs.PageIterate(func(instance tree.Page) bool {
		for _, source := range instance.GetSources() {
			r.parser.AddSource(source)
		}
		return false
	})
}

func (r *Runtime) Deal(sentence string) (concept.Index, error) {
	road, err := r.parser.Instance(sentence)

	if err != nil {
		return nil, err
	}

	roots := road.GetActiveSection()

	if len(roots) == 0 {
		return nil, errors.New(r.structErrorFormat(road))
	}
	if len(roots) != 1 {
		return nil, errors.New(r.priorityErrorFormat(road))
	}
	if roots[0].HasPriority() {
		return nil, errors.New(r.priorityErrorFormat(road))
	}

	return roots[0].Index(), nil
}

func (r *Runtime) Read(stream *os.File) error {
	var scanErr error = nil
	NewScan(&ScanParam{
		Stream: stream,
		OnReader: func(input string) bool {
			index, err := r.Deal(input)
			if err != nil {
				scanErr = err
				return false
			}
			r.Exec(index)
			return true
		},
		BeforeReader: func() {},
	}).Run()

	return scanErr
}

func (r *Runtime) Start() error {
	return r.box.Start()
}

func (r *Runtime) Stop() error {
	return r.box.Stop()
}

func (r *Runtime) Exec(hand concept.Index) {
	r.box.Exec(hand)
}

type RuntimeParam struct {
	OnError   func(error)
	OnPrint   func(concept.Variable)
	EventSize int
}

func NewRuntime(param *RuntimeParam) *Runtime {
	runtime := &Runtime{
		parser:              parser.NewParser(),
		libs:                NewLibraryManager(),
		structErrorFormat:   runtimeStructErrorFormatDefault,
		priorityErrorFormat: runtimePriorityErrorFormatDefault,
	}
	runtime.rootSpace = runtime.libs.Sandbox.Closure.New(nil)
	runtime.box = sandbox.NewSandbox(&sandbox.SandboxParam{
		Root:      runtime.rootSpace,
		OnError:   param.OnError,
		OnPrint:   param.OnPrint,
		EventSize: param.EventSize,
	})

	return runtime
}
