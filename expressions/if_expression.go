package expressions

import (
	"fmt"
	"strings"
)

type IfExpression struct {
	BooleanExpression Expression
	Block             Expression
}

func (v IfExpression) Execute(ctx *Context) (Value, error) {
	val, _ := v.BooleanExpression.Execute(ctx)

	if !val.IsBool {
		panic("if statement is not bool")
	}

	var lastValue Value
	if val.Bool {
		lastValue, _ = v.Block.Execute(ctx)
	}
	return lastValue, nil
}

func (v IfExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, "if block")
	v.BooleanExpression.Print(depth + 1)
	v.Block.Print(depth + 1)
}
