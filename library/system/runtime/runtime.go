package runtime

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type RuntimeParam struct {
	RootSpace concept.Pool
}

type Runtime struct {
	concept.Page
	param     *RuntimeParam
	libs      *tree.LibraryManager
	rootSpace concept.Object
}

func NewRuntime(libs *tree.LibraryManager, param *RuntimeParam) *Runtime {
	instance := &Runtime{
		libs:      libs,
		param:     param,
		Page:      libs.Sandbox.Variable.Page.New(),
		rootSpace: newPoolObject(libs, param.RootSpace),
	}

	instance.SetPublic(
		libs.Sandbox.Variable.String.New("rootSpace"),
		libs.Sandbox.Index.PublicIndex.New(
			"rootSpace",
			libs.Sandbox.Index.ConstIndex.New(instance.rootSpace),
		),
	)
	instance.SetPublic(
		libs.Sandbox.Variable.String.New("sandbox"),
		libs.Sandbox.Index.PublicIndex.New(
			"sandbox",
			libs.Sandbox.Index.ConstIndex.New(newSandboxCreator(libs)),
		),
	)

	return instance
}
