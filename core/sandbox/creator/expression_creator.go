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
	ClassRegister   *expression.ClassRegisterCreator
	For             *expression.ForCreator
	If              *expression.IfCreator
	FunctionEnd     *expression.FunctionEndCreator
	Return          *expression.ReturnCreator
	ParamSet        *expression.ParamSetCreator
	ParamGet        *expression.ParamGetCreator
	NewParam        *expression.NewParamCreator
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
}

func NewExpressionCreator(param *ExpressionCreatorParam) *ExpressionCreator {
	instance := &ExpressionCreator{}
	instance.NewParam = expression.NewNewParamCreator(&expression.NewParamCreatorParam{
		ParamCreator:           param.ParamCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.ParamGet = expression.NewParamGetCreator(&expression.ParamGetCreatorParam{
		ExceptionCreator:       param.ExceptionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.ParamSet = expression.NewParamSetCreator(&expression.ParamSetCreatorParam{
		ExceptionCreator:       param.ExceptionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.Return = expression.NewReturnCreator(&expression.ReturnCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.ExpressionIndex = adaptor.NewExpressionIndexCreator(&adaptor.ExpressionIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})
	instance.FunctionEnd = expression.NewFunctionEndCreator(&expression.FunctionEndCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		EndCreator:             param.EndCreator,
	})
	instance.If = expression.NewIfCreator(&expression.IfCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		CodeBlockCreator:       param.CodeBlockCreator,
		ExceptionCreator:       param.ExceptionCreator,
		ClosureCreator:         param.ClosureCreator,
	})
	instance.For = expression.NewForCreator(&expression.ForCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		StringCreator:          param.StringCreator,
		BoolCreator:            param.BoolCreator,
		CodeBlockCreator:       param.CodeBlockCreator,
		ExceptionCreator:       param.ExceptionCreator,
		ConstIndexCreator:      param.ConstIndexCreator,
	})
	instance.ClassRegister = expression.NewClassRegisterCreator(&expression.ClassRegisterCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		ExceptionCreator:       param.ExceptionCreator,
	})
	instance.Assignment = expression.NewAssignmentCreator(&expression.AssignmentCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.Call = expression.NewCallCreator(&expression.CallCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		ExceptionCreator:       param.ExceptionCreator,
		ParamCreator:           param.ParamCreator,
		ConstIndexCreator:      param.ConstIndexCreator,
	})
	return instance
}
