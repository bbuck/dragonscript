package lexer

const opString = "+-/*<>=!&|~"

func operatorState(l *Lexer) lexerState {
	l.run(opString)
	l.emit(TokenOperator)

	return whitespaceState
}
