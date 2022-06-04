package expressions

import (
	"fmt"
	"os"
)

type PrintLnFunction struct {
}

func (pl *PrintLnFunction) Execute(vals []Value) Value {
	args := make([]interface{}, 0)
	for _, val := range vals {
		if val.IsNumber {
			args = append(args, val.Number)
		} else if val.IsString {
			args = append(args, val.String)
		} else if val.IsBool {
			args = append(args, val.Bool)
		}
	}

	fmt.Println(args...)
	return Value{}
}

type ExitFunction struct {
}

func (ef *ExitFunction) Execute(val []Value) Value {
	if len(val) == 0 || len(val) > 1 {
		panic("exit accepts only 1 argument")
	}

	if !val[0].IsNumber {
		panic("exit function has to have a number as parameter")
	}
	os.Exit(val[0].Number)
	return Value{}
}
