package brackets

const (
	openSquare  = '['
	closeSquare = ']'
	openRound   = '('
	closeRound  = ')'
	openCurly   = '{'
	closeCurly  = '}'
)

func Bracket(input string) bool {
	stack := []rune{}

	for _, r := range input {
		switch r {
		case openSquare, openRound, openCurly:
			stack = push(stack, r)
		case closeSquare, closeRound, closeCurly:
			if !isMatch(stack, r) {
				return false
			}
			_, stack = pop(stack)
		}
	}

	return len(stack) == 0
}

func push(stack []rune, r rune) []rune {
	return append(stack, r)
}

func pop(stack []rune) (rune, []rune) {
	r := stack[len(stack)-1]
	s := stack[:len(stack)-1]
	return r, s
}

func peek(stack []rune) rune {
	return stack[len(stack)-1]
}

func isMatch(stack []rune, closing rune) bool {
	switch {
	case len(stack) == 0:
		return false
	case peek(stack) == openCurly && closing == closeCurly:
		return true
	case peek(stack) == openSquare && closing == closeSquare:
		return true
	case peek(stack) == openRound && closing == closeRound:
		return true
	default:
		return false
	}
}
