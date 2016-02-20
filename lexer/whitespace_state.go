package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

func whitespaceState(l *Lexer) lexerState {
	next := l.next()
	// we exclude newlines as space characters because they're terminators
	for next != '\n' && unicode.IsSpace(next) {
		next = l.next()
	}
	l.reverse()
	l.ignore()

	switch {
	case strings.ContainsRune(numberString, next):
		return intState
	case next == '\n' || next == ';':
		l.next()
		l.emit(TokenTerminator)

		// we handle moving to the next line here
		if next == '\n' {
			l.line++
			l.pos = 0
		}

		return whitespaceState
	case next == '.':
		return floatState
	case strings.ContainsRune(opString, next):
		return operatorState
	case next == EOF:
		return eofState
	case next == '(':
		l.next()
		l.emit(TokenLeftParenthesis)
	default:
		l.err = fmt.Errorf("unexpected character found at (%d:%d): %q", l.line, l.pos, next)
	}

	return nil
}
