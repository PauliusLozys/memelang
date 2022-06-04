package langparser

import (
	"fmt"
	"memelang/expressions"
	"strconv"
)

const enableLogging = false

func log(msg ...any) {
	if !enableLogging {
		return
	}
	fmt.Println(msg...)
}

type Parser struct {
	index int
	text  string
}

func NewTokenizer(text string) *Parser {
	return &Parser{
		text: text,
	}
}

func (t *Parser) ParseFile() expressions.Expression {
	statements := make([]expressions.Expression, 0)
	for t.index < len(t.text) {
		statements = append(statements, t.parseStatement())
	}
	return expressions.BlockExpression{
		Statements: statements,
	}
}

func (t *Parser) peekToken() string {
	// log("peekToken()")
	index := t.index
	for index < len(t.text) && (t.text[index] == ' ' || t.text[index] == '\r' || t.text[index] == '\n') { // skip whitespace
		index++
	}

	if index < len(t.text) && BreakOnParsing[t.text[index]] {
		return string(t.text[index])
	}

	token := []byte("")
	for index < len(t.text) && t.text[index] != ' ' && t.text[index] != '\r' {
		if t.text[index] == '\n' {
			index++
			continue
		}
		if BreakOnParsing[t.text[index]] {
			break
		}
		token = append(token, t.text[index])
		index++
	}

	return string(token)
}

func (t *Parser) getToken() string {
	// log("getToken()")
	for t.index < len(t.text) && (t.text[t.index] == ' ' || t.text[t.index] == '\r' || t.text[t.index] == '\n') { // skip whitespace
		t.index++
	}

	if t.index < len(t.text) && BreakOnParsing[t.text[t.index]] {
		ch := t.text[t.index]
		t.index++
		return string(ch)
	}

	token := []byte("")
	for t.index < len(t.text) && t.text[t.index] != ' ' && t.text[t.index] != '\r' {
		if t.text[t.index] == '\n' {
			t.index++
			continue
		}

		if BreakOnParsing[t.text[t.index]] {
			break
		}

		token = append(token, t.text[t.index])
		t.index++
	}
	log("parsed token", string(token), "Token length", len(token))
	return string(token)
}

func (t *Parser) parseStatement() expressions.Expression {
	log("parseStatement()")

	token := t.peekToken()
	switch token {
	case "let":
		return t.parseVariable()
	case "export":
		return t.parseExport()
	case "if":
		return t.parseIfBlock()
	case "{":
		return t.parseBlock()
	case "\"":
		return t.parseString()
	case "true", "false":
		return t.parseBooleanExpression()
	case "loop":
		return t.parseLoop()
	case "fun":
		return t.parseFunction()
	case "match":
		return t.parseMatch()
	case "return":
		return t.parseReturn()
	default:
		if _, err := strconv.Atoi(token); err == nil {
			return t.parseNumber()
		}
		if _, ok := MatchToken[token]; ok {
			return t.parseMatchStatement()
		}
		return t.parseIdentifier()
	}
}

func (t *Parser) parseVariable() expressions.Expression {
	log("parseVariable()")
	t.getToken() // Consume let
	varName := t.getToken()

	variable := expressions.VariableExpression{
		VariableName: varName,
	}

	if t.peekToken() == "=" {
		t.getToken() // consume "="
		assignment := expressions.AssignmentExpression{
			LHS: variable,
		}

		assignment.RHS = t.parseStatement()
		return assignment
	}

	return variable
}

func (t *Parser) parseIdentifier() expressions.Expression {
	log("parseIdentifier()")
	name := t.getToken()
	identExpression := expressions.IdentifierExpression{
		Name: name,
	}

	if t.peekToken() == "(" {
		funCall := expressions.FunctionCallExpression{
			FunctionName: name,
			Parameters:   t.parseFunctionArguments(),
		}
		return funCall
	}

	if t.peekToken() == "=" {
		t.getToken() // consume "="
		assignment := expressions.AssignmentExpression{
			LHS: expressions.VariableExpression{
				VariableName: name,
			},
			RHS: t.parseStatement(),
		}
		return assignment
	}

	if _, ok := BinaryOps[t.peekToken()]; ok {
		sign := t.getToken() // Consume +
		opExpression := expressions.BinaryOperatorExpression{
			LHS:  identExpression,
			Sign: sign,
			RHS:  t.parseStatement(),
		}
		return opExpression
	}
	return identExpression
}

func (t *Parser) parseNumber() expressions.Expression {
	log("parseNumber()")
	num, err := strconv.Atoi(t.getToken()) // consume number
	if err != nil {
		panic(err)
	}

	numExpression := expressions.NumericValueExpression{
		Number: num,
	}

	if _, ok := BinaryOps[t.peekToken()]; ok {
		sign := t.getToken() // Consume +
		opExpression := expressions.BinaryOperatorExpression{
			LHS:  numExpression,
			Sign: sign,
			RHS:  t.parseStatement(),
		}
		return &opExpression
	}

	return numExpression
}

func (t *Parser) parseBlock() expressions.Expression {
	log("parseBlock()")
	t.getToken() // Consume "{"

	statements := make([]expressions.Expression, 0)

	for t.index < len(t.text) {
		if t.peekToken() == "}" {
			t.getToken()
			break
		}
		statements = append(statements, t.parseStatement())
	}

	return expressions.BlockExpression{
		Statements: statements,
	}
}

func (t *Parser) parseExport() expressions.Expression {
	log("parseExport()")
	t.getToken() // Consume "export"
	functionName := t.getToken()

	return expressions.ExportExpression{
		ExportFunctionName: functionName,
	}
}

func (t *Parser) parseFunctionArguments() []expressions.Expression {
	log("parseFunctionArguments()")
	t.getToken() // Consume "("
	args := make([]expressions.Expression, 0)
	for {
		args = append(args, t.parseStatement())
		if t.peekToken() != "," {
			break
		}
		if t.getToken() != "," {
			panic("expected , in argument")
		}
	}
	t.getToken() // consume ")"
	return args
}

func (t *Parser) parseString() expressions.Expression {
	log("parseString()")
	t.getToken() // Consume "

	str := ""
	for t.text[t.index] != '"' {
		str += string(t.text[t.index])
		t.index++
	}
	t.index++ // Consume "

	return expressions.StringExpression{
		Content: str,
	}
}

func (t *Parser) parseIfBlock() expressions.Expression {
	log("parseIfBlock()")
	t.getToken() // Consume "if"
	f := t.parseStatement()

	if t.peekToken() != "{" {
		panic("expected if block to start")
	}

	return expressions.IfExpression{
		BooleanExpression: f,
		Block:             t.parseBlock(),
	}
}

func (t *Parser) parseBooleanExpression() expressions.Expression {
	log("parseBooleanExpression()")
	bo := t.getToken() // Consume true/false
	boolean, err := strconv.ParseBool(bo)
	if err != nil {
		panic(err)
	}

	return expressions.BooleanExpression{
		Bool: boolean,
	}
}

func (t *Parser) parseLoop() expressions.Expression {
	log("parseLoop()")
	t.getToken() // Consume "loop"
	f := t.parseStatement()

	if t.peekToken() != "{" {
		panic("expected if block to start")
	}

	return expressions.LoopExpression{
		BooleanExpression: f,
		Block:             t.parseBlock(),
	}
}

func (t *Parser) parseFunction() expressions.Expression {
	log("parseFunction()")
	t.getToken() // Consume "fun"

	functionName := t.getToken()

	if t.peekToken() != "(" {
		panic("expected function argument to start")
	}
	args := t.parseFunctionArguments()
	if t.peekToken() != "{" {
		panic("expected function block to start")
	}

	block := t.parseBlock()

	return expressions.FunctionDeclarationExpression{
		FunctionName: functionName,
		Argument:     args,
		Block:        block,
	}
}

func (t *Parser) parseMatch() expressions.Expression {
	log("parseMatch()")
	t.getToken() // Consume "match"

	if t.peekToken() != "(" {
		panic("expected function argument to start")
	}

	args := t.parseFunctionArguments()
	if len(args) != 1 {
		panic("bad match expression argument count")
	}

	if t.peekToken() != "{" {
		panic("expected match block to start")
	}

	block := t.parseBlock()

	return expressions.MatchExpression{
		Parameter: args[0],
		Block:     block,
	}
}

func (t *Parser) parseMatchStatement() expressions.Expression {
	log("parseMatchStatement()")
	sign := t.getToken()    // Consume bool sign
	v := t.parseStatement() // consume value
	if t.getToken() != "=>" {
		panic("expected => in match statement")
	}

	args := t.parseStatement()

	if t.getToken() != "," {
		panic("unexpected match statement ending")
	}

	return expressions.MatchExpressionStatement{
		CompareSign: sign,
		Final:       args,
		CompareTo:   v,
	}
}

func (t *Parser) parseReturn() expressions.Expression {
	log("parseReturn()")
	t.getToken() // Consume "return"
	// consume value
	return expressions.ReturnStatement{
		Argument: t.parseStatement(),
	}
}
