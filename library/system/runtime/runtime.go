package runtime

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type RuntimeParam struct {
	RootSpace     concept.Pool
	RootPipeCache *tree.PipeCache
}

type Runtime struct {
	concept.Page
	param         *RuntimeParam
	libs          *tree.LibraryManager
	rootSpace     concept.Pool
	rootPipeCache *tree.PipeCache
}

func NewRuntime(libs *tree.LibraryManager, param *RuntimeParam) *Runtime {
	instance := &Runtime{
		libs:          libs,
		param:         param,
		Page:          libs.Sandbox.Variable.Page.New(),
		rootSpace:     param.RootSpace,
		rootPipeCache: param.RootPipeCache,
	}
	instance.fieldInit()
	instance.SetPublic(
		libs.Sandbox.Variable.DelayString.New("rootSpace"),
		libs.Sandbox.Index.PublicIndex.New(
			"rootSpace",
			libs.Sandbox.Index.ConstIndex.New(instance.rootSpace),
		),
	)
	instance.SetPublic(
		libs.Sandbox.Variable.DelayString.New("findHistory"),
		libs.Sandbox.Index.PublicIndex.New(
			"findHistory",
			libs.Sandbox.Index.ConstIndex.New(newFindHistory(libs, instance.rootSpace)),
		),
	)
	instance.SetPublic(
		libs.Sandbox.Variable.DelayString.New("findPipeCache"),
		libs.Sandbox.Index.PublicIndex.New(
			"findPipeCache",
			libs.Sandbox.Index.ConstIndex.New(newFindPipeCache(libs, instance.rootPipeCache)),
		),
	)
	instance.SetPublic(
		libs.Sandbox.Variable.DelayString.New("throw"),
		libs.Sandbox.Index.PublicIndex.New(
			"throw",
			libs.Sandbox.Index.ConstIndex.New(newThrow(libs)),
		),
	)

	return instance
}

func (r *Runtime) fieldInit() {
	r.libs.Sandbox.Variable.Pool.Inits = append(r.libs.Sandbox.Variable.Pool.Inits, func(instance *variable.Pool) {
		VariableHomeInit(r.libs, instance)
		PoolInit(r.libs, instance)
	})
	r.libs.Sandbox.Variable.Array.Inits = append(r.libs.Sandbox.Variable.Array.Inits, func(instance *variable.Array) {
		VariableHomeInit(r.libs, instance)
		ArrayInit(r.libs, instance)
	})
	r.libs.Sandbox.Variable.String.Inits = append(r.libs.Sandbox.Variable.String.Inits, func(instance *variable.String) {
		VariableHomeInit(r.libs, instance)
		StringInit(r.libs, instance)
	})
	r.libs.Sandbox.Variable.DefineFunction.Inits = append(r.libs.Sandbox.Variable.DefineFunction.Inits, func(instance *variable.DefineFunction) {
		VariableHomeInit(r.libs, instance)
		FunctionHomeInit(r.libs, instance)
	})
	r.libs.Sandbox.Variable.SystemFunction.Inits = append(r.libs.Sandbox.Variable.SystemFunction.Inits, func(instance *variable.SystemFunction) {
		VariableHomeInit(r.libs, instance)
		FunctionHomeInit(r.libs, instance)
	})
	r.libs.Sandbox.Variable.Function.Inits = append(r.libs.Sandbox.Variable.Function.Inits, func(instance *variable.Function) {
		VariableHomeInit(r.libs, instance)
		FunctionHomeInit(r.libs, instance)
	})
}
