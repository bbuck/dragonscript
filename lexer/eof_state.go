package lexer

func eofState(l *Lexer) lexerState {
	l.ignore()
	l.emit(TokenEOF)

	return nil
}
