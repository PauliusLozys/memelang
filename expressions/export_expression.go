package expressions

import (
	"fmt"
	"strings"
)

var ExportedFunctionMap = map[string]ExportedFunction{
	"println": &PrintLnFunction{},
	"exit":    &ExitFunction{},
}

type ExportExpression struct {
	ExportFunctionName string
}

func (i ExportExpression) Execute(ctx *Context) (Value, error) {
	//FIXME: have a lookup table based on the exported function name

	fn, ok := ExportedFunctionMap[i.ExportFunctionName]
	if !ok {
		panic("exported function does not exist")
	}

	ctx.ExportedFunctions[i.ExportFunctionName] = fn
	return Value{}, nil
}

func (v ExportExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, "export")
	fmt.Println(pad, v.ExportFunctionName)
}
