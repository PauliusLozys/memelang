package expressions

import (
	"errors"
	"fmt"
	"strings"
)

type FunctionCallExpression struct {
	FunctionName string
	Parameters   []Expression
}

func (i FunctionCallExpression) Execute(ctx *Context) (Value, error) {
	if fn, ok := ctx.ExportedFunctions[i.FunctionName]; ok {
		args := make([]Value, 0)
		for _, argument := range i.Parameters {
			arg, _ := argument.Execute(ctx)
			args = append(args, arg)
		}
		val := fn.Execute(args)
		return val, nil
	}
	if fn, ok := ctx.DeclaredFunctions[i.FunctionName]; ok {
		args := make([]Value, 0)
		for _, argument := range i.Parameters {
			arg, _ := argument.Execute(ctx)
			args = append(args, arg)
		}

		if len(args) != len(ctx.FunctionVariables[i.FunctionName]) {
			panic("mismatching declared and called function argument count")
		}

		ii := 0
		for v, _ := range ctx.FunctionVariables[i.FunctionName] {
			ctx.FunctionVariables[i.FunctionName][v] = args[ii]
			ii++
		}
		ctx.InFunction = true
		ctx.FunctionName = i.FunctionName

		val, _ := fn.Execute(ctx)

		ctx.InFunction = false
		ctx.FunctionName = ""

		return val, nil
	}

	panic(errors.New("unknown function call"))
}

func (v FunctionCallExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.FunctionName)
	fmt.Println(pad, v.Parameters)
}
