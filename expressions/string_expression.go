package expressions

import (
	"fmt"
	"strings"
)

type StringExpression struct {
	Content string
}

func (v StringExpression) Execute(ctx *Context) (Value, error) {
	return Value{
		String:   v.Content,
		IsString: true,
	}, nil
}

func (v StringExpression) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Println(pad, v.Content)
}
