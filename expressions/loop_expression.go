package expressions

type LoopExpression struct {
	BooleanExpression Expression
	Block             Expression
}

func (v LoopExpression) Execute(ctx *Context) (Value, error) {
	val, _ := v.BooleanExpression.Execute(ctx)

	if !val.IsBool {
		panic("if statement is not bool")
	}

	for val.Bool {
		if val.Bool {
			v.Block.Execute(ctx)
		}
		val, _ = v.BooleanExpression.Execute(ctx)
	}

	return Value{}, nil
}

func (v LoopExpression) Print(depth int) {
	v.BooleanExpression.Print(depth + 1)
	v.Block.Print(depth + 1)
}
