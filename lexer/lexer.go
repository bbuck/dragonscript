package lexer

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

// EOF is a rune that represents the end of the input.
const EOF rune = -1

type lexerState func(*Lexer) lexerState

// Lexer provides a means to browse an io.Reader and convert it's contents into
// a tokenized list that can then be put through some other process in the
// future to achieve a specific outcome
type Lexer struct {
	Tokens    []*Token
	src       *bytes.Buffer
	cur       *bytes.Buffer
	pos       int
	line      int
	prevSize  int
	tokenized bool
	err       error
}

// NewLexer takes an input source and returns a lexer that is ready to tokenize
// the source.
func NewLexer(src interface{}) (*Lexer, error) {
	var buf *bytes.Buffer
	switch srcVal := src.(type) {
	case string:
		buf = bytes.NewBufferString(srcVal)
	case *string:
		buf = bytes.NewBufferString(*srcVal)
	default:
		if rdr, ok := src.(io.Reader); ok {
			buf = new(bytes.Buffer)
			buf.ReadFrom(rdr)
		} else {
			return nil, errors.New("Invalid source provided to the lexer!")
		}
	}

	return &Lexer{
		Tokens:   make([]*Token, 0),
		src:      buf,
		cur:      new(bytes.Buffer),
		line:     1,
		prevSize: -1,
	}, nil
}

// Tokenize is a shortcut method that creates a lexer and then initiates the
// lexical analysis process and returns the tokens (or error) generated from
// the process. This is the primary way you should interact with lexers as they
// do not leave past a lexing process
func Tokenize(src interface{}) ([]*Token, error) {
	lex, err := NewLexer(src)
	if err != nil {
		return nil, err
	}
	err = lex.Analyze()
	if err != nil {
		return nil, err
	}

	return lex.Tokens, nil
}

// Analyze begins the tokenization process. Once start returns, provided no error
// is returned, then .Tokens will contain the list of tokens from the given input
// source.
func (l *Lexer) Analyze() error {
	state := whitespaceState

	for state != nil {
		state = state(l)
		if l.err != nil {
			return l.err
		}
	}

	return nil
}

func (l *Lexer) next() rune {
	if l.tokenized {
		return EOF
	}

	r, size, err := l.src.ReadRune()
	if err != nil {
		if err != io.EOF {
			l.err = err
		}

		l.tokenized = true
		return EOF
	}
	l.prevSize = size
	l.cur.WriteRune(r)

	l.pos++

	return r
}

func (l *Lexer) ignore() {
	l.cur.Reset()
	l.prevSize = -1
}

func (l *Lexer) run(valid string) rune {
	next := l.next()
	for strings.ContainsRune(valid, next) {
		next = l.next()
	}
	l.reverse()

	return next
}

func (l *Lexer) emit(t TokenType) {
	val := l.cur.String()
	l.cur.Reset()
	l.prevSize = -1
	l.Tokens = append(l.Tokens, &Token{Type: t, Value: val})
}

func (l *Lexer) reverse() {
	if l.prevSize < 0 || l.tokenized {
		return
	}

	err := l.src.UnreadRune()
	if err != nil && l.err == nil {
		l.err = err
	}
	l.cur.Truncate(l.cur.Len() - l.prevSize)
	l.prevSize = -1

	l.pos--
}

func (l *Lexer) peek() rune {
	r := l.next()
	l.reverse()

	return r
}
