package main

import (
	"fmt"
	"memelang/expressions"
	"memelang/langparser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please provide a source file")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tokenizer := langparser.NewTokenizer(string(data))
	ctx := expressions.NewContext()
	block := tokenizer.ParseFile()
	// block.Print(0)

	block.Execute(ctx)
	//fmt.Println(ctx.LocalVariables)
}
