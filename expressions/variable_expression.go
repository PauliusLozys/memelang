package expressions

import (
	"fmt"
	"strings"
)

type VariableExpression struct {
	VariableName string
}

func (v VariableExpression) Execute(ctx *Context) (Value, error) {
	return Value{
		String:     v.VariableName,
		IsVariable: true,
	}, nil
}

func (v VariableExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.VariableName)
}
