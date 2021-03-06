package runtime

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler"
	"github.com/TingerSure/natural_language/core/parser"
	"github.com/TingerSure/natural_language/core/sandbox"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/creator"
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	runtimeStructErrorFormatDefault = func(road *parser.Road) string {
		maxMatch := ""
		for index := road.SentenceSize() - 1; index >= 0; {
			section := road.GetSectionMax(index)
			maxMatch = section.ToString() + maxMatch
			index -= section.ContentSize()
		}
		return fmt.Sprintf("%v\nNo struct rule can match this sentence.", maxMatch)
	}

	runtimePriorityErrorFormatDefault = func(road *parser.Road) string {
		roots := road.GetActiveSection()
		if len(roots) == 1 {
			return fmt.Sprintf("%v \nNo priority rule can distinguish all meanings.", roots[0].ToString())
		}
		return fmt.Sprintf("%v \nNo priority rule can distinguish all meanings.", tree.NewPhrasePriority(roots).ToString())
	}
)

type Runtime struct {
	parser              *parser.Parser
	compiler            *compiler.Compiler
	libs                *tree.LibraryManager
	box                 *sandbox.Sandbox
	rootPipeCache       *tree.PipeCache
	rootSpace           concept.Pool
	defaultLanguage     string
	structErrorFormat   func(*parser.Road) string
	priorityErrorFormat func(*parser.Road) string
}

func (r *Runtime) GetRootPipeCache() *tree.PipeCache {
	return r.rootPipeCache
}

func (r *Runtime) GetRootSpace() concept.Pool {
	return r.rootSpace
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

func (r *Runtime) GetLibraryManager() *tree.LibraryManager {
	return r.libs
}

func (r *Runtime) Deal(sentence string) (concept.Function, error) {
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

	return roots[0].Index()
}

func (r *Runtime) Read(path string) error {
	return r.compiler.Read(path)
}

func (r *Runtime) Start() error {
	return r.box.Start()
}

func (r *Runtime) Stop() error {
	return r.box.Stop()
}

func (r *Runtime) Exec(hand concept.Function, line concept.Line) {
	r.box.Exec(hand, line)
}

type RuntimeParam struct {
	OnError         func(error)
	OnPrint         func(concept.Variable)
	EventSize       int
	SourceRoots     []string
	SourceExtension string
}

func NewRuntime(param *RuntimeParam) *Runtime {
	sandboxCreator := creator.NewSandboxCreator()
	runtime := &Runtime{
		structErrorFormat:   runtimeStructErrorFormatDefault,
		priorityErrorFormat: runtimePriorityErrorFormatDefault,
	}
	runtime.rootPipeCache = tree.NewPipeCache(&tree.PipeCacheParam{})
	runtime.rootSpace = sandboxCreator.Variable.Pool.New(nil)
	runtime.parser = parser.NewParser(runtime.rootSpace, sandboxCreator)
	runtime.box = sandbox.NewSandbox(&sandbox.SandboxParam{
		Root:    runtime.rootSpace,
		OnError: param.OnError,
		OnPrint: param.OnPrint,
		OnExec: func(pipe concept.Function, value concept.Variable) {
			runtime.rootPipeCache.Add(pipe, value)
		},
		EventSize: param.EventSize,
	})
	runtime.libs = tree.NewLibraryManager(
		sandboxCreator,
		runtime.parser.GetLexer(),
		runtime.parser.GetReach(),
		runtime.parser.GetBarricade(),
		runtime.parser.GetTypes(),
		runtime.parser.GetDiversion(),
	)
	runtime.compiler = compiler.NewCompiler(runtime.libs)
	runtime.compiler.AddRoots(param.SourceRoots...)
	runtime.compiler.SetExtension(param.SourceExtension)
	return runtime
}
