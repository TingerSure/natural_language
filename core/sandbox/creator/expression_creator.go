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
	ExpressionIndex   *adaptor.ExpressionIndexCreator
	Call              *expression.CallCreator
	Assignment        *expression.AssignmentCreator
	For               *expression.ForCreator
	If                *expression.IfCreator
	Define            *expression.DefineCreator
	NewParam          *expression.NewParamCreator
	Component         *expression.ComponentCreator
	NewString         *expression.NewStringCreator
	NewBool           *expression.NewBoolCreator
	NewNumber         *expression.NewNumberCreator
	NewObject         *expression.NewObjectCreator
	NewNull           *expression.NewNullCreator
	NewFunction       *expression.NewFunctionCreator
	NewBreak          *expression.NewBreakCreator
	NewContinue       *expression.NewContinueCreator
	NewReturn         *expression.NewReturnCreator
	NewDefineFunction *expression.NewDefineFunctionCreator
	NewClass          *expression.NewClassCreator
	NewMappingObject  *expression.NewMappingObjectCreator
	Parenthesis       *expression.ParenthesisCreator
	NewArray          *expression.NewArrayCreator
	IndexComponent    *expression.IndexComponentCreator
	Append            *expression.AppendCreator
}

type ExpressionCreatorParam struct {
	CodeBlockCreator      func() *code_block.CodeBlock
	DefineFunctionCreator func([]concept.String, []concept.String) *variable.DefineFunction
	FunctionCreator       func(concept.Pool) *variable.Function
	StringCreator         func(string) concept.String
	BoolCreator           func(bool) concept.Bool
	NumberCreator         func(float64) concept.Number
	ExceptionCreator      func(string, string) concept.Exception
	ReturnCreator         func() *interrupt.Return
	ParamCreator          func() concept.Param
	ConstIndexCreator     func(concept.Variable) *index.ConstIndex
	PoolCreator           func(concept.Pool) concept.Pool
	NullCreator           func() concept.Null
	ObjectCreator         func() concept.Object
	MappingObjectCreator  func(concept.Variable, concept.Class) *variable.MappingObject
	ContinueCreator       func(concept.String) *interrupt.Continue
	BreakCreator          func(concept.String) *interrupt.Break
	ClassCreator          func() concept.Class
	ArrayCreator          func() *variable.Array
}

func NewExpressionCreator(param *ExpressionCreatorParam) *ExpressionCreator {
	instance := &ExpressionCreator{}
	instance.ExpressionIndex = adaptor.NewExpressionIndexCreator(&adaptor.ExpressionIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
	})
	instance.Append = expression.NewAppendCreator(&expression.AppendCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		ExceptionCreator:       param.ExceptionCreator,
	})
	instance.IndexComponent = expression.NewIndexComponentCreator(&expression.IndexComponentCreatorParam{
		ExceptionCreator:       param.ExceptionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewArray = expression.NewNewArrayCreator(&expression.NewArrayCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		ArrayCreator:           param.ArrayCreator,
	})
	instance.Parenthesis = expression.NewParenthesisCreator(&expression.ParenthesisCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewMappingObject = expression.NewNewMappingObjectCreator(&expression.NewMappingObjectCreatorParam{
		MappingObjectCreator:   param.MappingObjectCreator,
		ExceptionCreator:       param.ExceptionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewClass = expression.NewNewClassCreator(&expression.NewClassCreatorParam{
		ClassCreator:           param.ClassCreator,
		StringCreator:          param.StringCreator,
		ExceptionCreator:       param.ExceptionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewDefineFunction = expression.NewNewDefineFunctionCreator(&expression.NewDefineFunctionCreatorParam{
		DefineFunctionCreator:  param.DefineFunctionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewReturn = expression.NewNewReturnCreator(&expression.NewReturnCreatorParam{
		ReturnCreator:          param.ReturnCreator,
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewBreak = expression.NewNewBreakCreator(&expression.NewBreakCreatorParam{
		BreakCreator:           param.BreakCreator,
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewContinue = expression.NewNewContinueCreator(&expression.NewContinueCreatorParam{
		ContinueCreator:        param.ContinueCreator,
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewFunction = expression.NewNewFunctionCreator(&expression.NewFunctionCreatorParam{
		FunctionCreator:        param.FunctionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewNull = expression.NewNewNullCreator(&expression.NewNullCreatorParam{
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewObject = expression.NewNewObjectCreator(&expression.NewObjectCreatorParam{
		ObjectCreator:          param.ObjectCreator,
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewNumber = expression.NewNewNumberCreator(&expression.NewNumberCreatorParam{
		NumberCreator:          param.NumberCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewBool = expression.NewNewBoolCreator(&expression.NewBoolCreatorParam{
		BoolCreator:            param.BoolCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewString = expression.NewNewStringCreator(&expression.NewStringCreatorParam{
		StringCreator:          param.StringCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.Component = expression.NewComponentCreator(&expression.ComponentCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.NewParam = expression.NewNewParamCreator(&expression.NewParamCreatorParam{
		ParamCreator:           param.ParamCreator,
		ExceptionCreator:       param.ExceptionCreator,
		NullCreator:            param.NullCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.Define = expression.NewDefineCreator(&expression.DefineCreatorParam{
		NullCreator:            param.NullCreator,
		ExceptionCreator:       param.ExceptionCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
	})
	instance.If = expression.NewIfCreator(&expression.IfCreatorParam{
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		CodeBlockCreator:       param.CodeBlockCreator,
		ExceptionCreator:       param.ExceptionCreator,
		NullCreator:            param.NullCreator,
		PoolCreator:            param.PoolCreator,
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
		StringCreator:          param.StringCreator,
		ExpressionIndexCreator: instance.ExpressionIndex.New,
		ExceptionCreator:       param.ExceptionCreator,
		ParamCreator:           param.ParamCreator,
		ConstIndexCreator:      param.ConstIndexCreator,
		NewParamCreator:        instance.NewParam.New,
	})
	return instance
}
