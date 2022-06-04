package expressions

type Context struct {
	LocalVariables    map[string]Value
	ExportedFunctions map[string]ExportedFunction
	FunctionVariables map[string]map[string]Value
	DeclaredFunctions map[string]Expression
	InFunction        bool
	InMatch           bool
	MatchArgument     Expression
	FunctionName      string
}

func NewContext() *Context {
	return &Context{
		LocalVariables:    make(map[string]Value),
		ExportedFunctions: make(map[string]ExportedFunction),
		FunctionVariables: make(map[string]map[string]Value),
		DeclaredFunctions: make(map[string]Expression),
	}
}
