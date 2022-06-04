package expressions

import (
	"fmt"
	"strings"
)

type MatchExpressionStatement struct {
	Final       Expression
	CompareSign string
	CompareTo   Expression
}

func (v MatchExpressionStatement) Execute(ctx *Context) (Value, error) {
	if !ctx.InMatch {
		panic("cant be in match?")
	}
	exp := BinaryOperatorExpression{
		LHS:  ctx.MatchArgument,
		Sign: v.CompareSign,
		RHS:  v.CompareTo,
	}

	val, _ := exp.Execute(ctx)
	if !val.IsBool {
		panic("match should have boolean comparisons")
	}

	if val.Bool {
		v, _ := v.Final.Execute(ctx)
		v.BreakBlock = true
		return v, nil
	}

	return Value{
		Number:   -1,
		IsNumber: true,
	}, nil
}

func (v MatchExpressionStatement) Print(depth int) {
	v.Final.Print(depth + 1)
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.CompareSign)
	v.CompareTo.Print(depth + 1)
}
