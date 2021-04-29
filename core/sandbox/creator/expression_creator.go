package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ExpressionCreator struct {
	ExpressionIndex *adaptor.ExpressionIndexCreator
	Call            *expression.CallCreator
	Assignment      *expression.AssignmentCreator
	For             *expression.ForCreator
	If              *expression.IfCreator
	FunctionEnd     *expression.FunctionEndCreator
	Return          *expression.ReturnCreator
	NewParam        *expression.NewParamCreator
	Component       *expression.ComponentCreator
}

type ExpressionCreatorParam struct {
	CodeBlockCreator  func() *code_block.CodeBlock
	StringCreator     func(string) concept.String
	BoolCreator       func(bool) *variable.Bool
	ExceptionCreator  func(string, string) concept.Exception
	EndCreator        func() *interrupt.End
	ParamCreator      func() concept.Param
	ConstIndexCreator func(concept.Variable) *index.ConstIndex
	ClosureCreator    func(concept.Closure) concept.Closure
	NullCreator       func() concept.Null
}

func NewExpressionCreator(param *ExpressionCreatorParam) *ExpressionCreator {
	instance := &ExpressionCreator{}
	instance.Component = expression.NewComponentCreator(&expression.ComponentCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewParam = expression.NewNewParamCreator(&expression.NewParamCreatorParam{
		ParamCreator:           param.ParamCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.Return = expression.NewReturnCreator(&expression.ReturnCreatorParam{
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.ExpressionIndex = adaptor.NewExpressionIndexCreator(&adaptor.ExpressionIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
	})
	instance.FunctionEnd = expression.NewFunctionEndCreator(&expression.FunctionEndCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		EndCreator:             param.EndCreator,
		NullCreator:            param.NullCreator,
	})
	instance.If = expression.NewIfCreator(&expression.IfCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		CodeBlockCreator:       param.CodeBlockCreator,
		ExceptionCreator:       param.ExceptionCreator,
		NullCreator:            param.NullCreator,
		ClosureCreator:         param.ClosureCreator,
	})
	instance.For = expression.NewForCreator(&expression.ForCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		StringCreator:          param.StringCreator,
		BoolCreator:            param.BoolCreator,
		CodeBlockCreator:       param.CodeBlockCreator,
		ExceptionCreator:       param.ExceptionCreator,
		ConstIndexCreator:      param.ConstIndexCreator,
		NullCreator:            param.NullCreator,
	})
	instance.Assignment = expression.NewAssignmentCreator(&expression.AssignmentCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.Call = expression.NewCallCreator(&expression.CallCreatorParam{
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		ExceptionCreator:       param.ExceptionCreator,
		ParamCreator:           param.ParamCreator,
		ConstIndexCreator:      param.ConstIndexCreator,
	})
	return instance
}
