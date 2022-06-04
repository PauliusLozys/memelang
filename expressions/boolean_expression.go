package expressions

import (
	"fmt"
	"strings"
)

type BooleanExpression struct {
	Bool bool
}

func (v BooleanExpression) Execute(ctx *Context) (Value, error) {
	return Value{
		Bool:   v.Bool,
		IsBool: true,
	}, nil
}

func (v BooleanExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.Bool)
}
