package expressions

// General value for all expressions
type Value struct {
	Number     int
	String     string
	Type       string
	Bool       bool
	IsVariable bool
	IsNumber   bool
	IsString   bool
	IsFunction bool
	IsBool     bool
	BreakBlock bool
}

type Expression interface {
	// Execute executes expression statement and its result is returned.
	Execute(ctx *Context) (Value, error)
	// Prints rough AST expression
	Print(depth int)
}

type ExportedFunction interface {
	// Executes exported function with Value argument.
	Execute(val []Value) Value
}
