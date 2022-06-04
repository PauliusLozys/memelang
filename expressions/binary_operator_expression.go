package expressions

import (
	"errors"
	"fmt"
	"strings"
)

type BinaryOperatorExpression struct {
	LHS  Expression
	RHS  Expression
	Sign string
}

func (bo BinaryOperatorExpression) Execute(ctx *Context) (Value, error) {

	val1, _ := bo.LHS.Execute(ctx)
	val2, _ := bo.RHS.Execute(ctx)

	if bo.Sign == "+" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsNumber: true,
			Number:   val1.Number + val2.Number,
		}, nil
	}
	if bo.Sign == "-" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsNumber: true,
			Number:   val1.Number - val2.Number,
		}, nil
	}
	if bo.Sign == "/" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsNumber: true,
			Number:   val1.Number / val2.Number,
		}, nil
	}
	if bo.Sign == "*" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsNumber: true,
			Number:   val1.Number * val2.Number,
		}, nil
	}
	if bo.Sign == "&&" && val1.IsBool && val2.IsBool {
		return Value{
			IsBool: true,
			Bool:   val1.Bool && val2.Bool,
		}, nil
	}
	if bo.Sign == "||" && val1.IsBool && val2.IsBool {
		return Value{
			IsBool: true,
			Bool:   val1.Bool || val2.Bool,
		}, nil
	}
	if bo.Sign == "==" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsBool: true,
			Bool:   val1.Number == val2.Number,
		}, nil
	}
	if bo.Sign == "==" && val1.IsBool && val2.IsBool {
		return Value{
			IsBool: true,
			Bool:   val1.Bool == val2.Bool,
		}, nil
	}
	if bo.Sign == "!=" && val1.IsBool && val2.IsBool {
		return Value{
			IsBool: true,
			Bool:   val1.Bool != val2.Bool,
		}, nil
	}
	if bo.Sign == "!=" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsBool: true,
			Bool:   val1.Number != val2.Number,
		}, nil
	}
	if bo.Sign == "<" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsBool: true,
			Bool:   val1.Number < val2.Number,
		}, nil
	}
	if bo.Sign == ">" && val1.IsNumber && val2.IsNumber {
		return Value{
			IsBool: true,
			Bool:   val1.Number > val2.Number,
		}, nil
	}

	fmt.Println(val1)
	fmt.Println(val2)

	panic(errors.New("unknown in BinaryOperatorExpression Execute()"))
}

func (v BinaryOperatorExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	v.LHS.Print(depth + 1)
	fmt.Println(pad, v.Sign)
	v.RHS.Print(depth + 1)
}
