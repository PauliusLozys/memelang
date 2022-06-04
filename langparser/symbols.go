package langparser

var (
	BinaryOps = map[string]struct{}{
		"+":  {},
		"-":  {},
		"/":  {},
		"*":  {},
		"&&": {},
		"||": {},
		"==": {},
		"!=": {},
		">":  {},
		"<":  {},
	}
	BreakOnParsing = map[byte]bool{
		'(': true,
		')': true,
		'{': true,
		'}': true,
		'+': true,
		'-': true,
		'"': true,
		',': true,
	}
	MatchToken = map[string]struct{}{
		"==": {},
		"!=": {},
		">":  {},
		"<":  {},
	}
)
