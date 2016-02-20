package lexer

const numberString = "0123456789"

func intState(l *Lexer) lexerState {
	numbers := numberString + "_"
	l.next()
	if l.peek() == 'x' || l.peek() == 'X' {
		l.next()
		numbers += "abcdefABCDEF"
	}
	final := l.run(numbers)

	if final == '.' {
		return floatState
	}

	l.emit(TokenInt)

	return whitespaceState
}

func floatState(l *Lexer) lexerState {
	if l.next() != '.' {
		return whitespaceState
	}

	l.run(numberString + "_")
	l.emit(TokenFloat)

	return whitespaceState
}
