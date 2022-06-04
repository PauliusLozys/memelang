package expressions

import (
	"fmt"
	"strings"
)

type MatchExpression struct {
	Parameter Expression
	Block     Expression
}

func (v MatchExpression) Execute(ctx *Context) (Value, error) {
	ctx.InMatch = true
	ctx.MatchArgument = v.Parameter

	val, _ := v.Block.Execute(ctx)

	return val, nil
}

func (v MatchExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, "match")
	v.Parameter.Print(depth + 1)
	v.Block.Print(depth + 1)
}
