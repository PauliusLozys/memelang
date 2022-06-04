package expressions

import (
	"fmt"
	"strings"
)

type FunctionDeclarationExpression struct {
	FunctionName string
	Argument     []Expression
	Block        Expression
}

func (i FunctionDeclarationExpression) Execute(ctx *Context) (Value, error) {
	localMap := make(map[string]Value)
	for _, argument := range i.Argument {
		arg, _ := argument.Execute(ctx)
		if !arg.IsVariable {
			panic("function declaration expected a variable")
		}
		localMap[arg.String] = Value{}
	}
	ctx.FunctionVariables[i.FunctionName] = localMap
	ctx.DeclaredFunctions[i.FunctionName] = i.Block
	return Value{}, nil
}

func (v FunctionDeclarationExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.FunctionName)
	//FIXME: print properly
	// v.Argument.Print(depth + 1)
	v.Block.Print(depth + 1)
}
