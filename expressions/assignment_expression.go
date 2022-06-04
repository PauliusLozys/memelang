package expressions

import (
	"fmt"
	"strings"
)

type AssignmentExpression struct {
	LHS Expression
	RHS Expression
}

func (v AssignmentExpression) Execute(ctx *Context) (Value, error) {
	val1, _ := v.LHS.Execute(ctx)
	val2, _ := v.RHS.Execute(ctx)
	val2.Type = val1.Type

	if !val1.IsVariable {
		panic("assigned value is not a variable")
	}

	ctx.LocalVariables[val1.String] = val2
	return Value{}, nil
}

func (v AssignmentExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	v.LHS.Print(depth + 1)
	fmt.Println(pad, "=")
	v.RHS.Print(depth + 1)
}
