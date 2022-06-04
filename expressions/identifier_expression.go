package expressions

import (
	"fmt"
	"strings"
)

type IdentifierExpression struct {
	Name string
}

func (i IdentifierExpression) Execute(ctx *Context) (Value, error) {
	if val, ok := ctx.FunctionVariables[ctx.FunctionName][i.Name]; ctx.InFunction && ok {
		return val, nil
	}

	if val, ok := ctx.LocalVariables[i.Name]; ok {
		return val, nil
	}
	panic("unknown function variable")
}

func (v IdentifierExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.Name)
}
