package expressions

import (
	"fmt"
	"strings"
)

type NumericValueExpression struct {
	Number int
}

func (v NumericValueExpression) Execute(ctx *Context) (Value, error) {
	return Value{
		Number:   v.Number,
		IsNumber: true,
	}, nil
}

func (v NumericValueExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.Number)
}
