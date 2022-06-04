package expressions

type ReturnStatement struct {
	Argument Expression
}

func (v ReturnStatement) Execute(ctx *Context) (Value, error) {
	val, _ := v.Argument.Execute(ctx)
	val.BreakBlock = true
	return val, nil
}

func (v ReturnStatement) Print(depth int) {
	v.Argument.Print(depth + 1)
}
