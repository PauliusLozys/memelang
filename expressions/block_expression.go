package expressions

import (
	"fmt"
	"strings"
)

type BlockExpression struct {
	Statements []Expression
}

func (b BlockExpression) Execute(ctx *Context) (Value, error) {
	var lastValue Value
	for _, statement := range b.Statements {
		lastValue, _ = statement.Execute(ctx)
		if lastValue.BreakBlock {
			break
		}
	}
	return lastValue, nil
}

func (b BlockExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, "Block statement")
	for _, statement := range b.Statements {
		statement.Print(depth + 1)
		fmt.Println("=============================")
	}
}
