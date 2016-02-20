package lexer

import "fmt"

// TokenType represents the type the value will be. This can be a native Go
// type or a deeper nested Go struct that fully represents the data the token
// contains.
type TokenType uint8

// Various types of tokens supported by this lexer
const (
	TokenEOF TokenType = iota
	TokenInt
	TokenFloat
	TokenTerminator
	TokenOperator
	TokenRightParenthesis
	TokenLeftParenthesis
	TokenRightBrace
	TokenLeftBrace
	TokenRightBracket
	TokenLeftBracket
)

func (t TokenType) String() string {
	switch t {
	case TokenEOF:
		return "EOF"
	case TokenInt:
		return "int"
	case TokenFloat:
		return "float"
	case TokenTerminator:
		return "term"
	case TokenOperator:
		return "op"
	}

	return "WUT?"
}

// Token is a struct that represents a tokenized value. The value's type is
// dependent on what type of token it is. For example, if the token type is
// TokenInt then the value is most likely an int64 value.
type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	if t.Type == TokenEOF {
		return "[ <<EOF>> ]"
	}

	return fmt.Sprintf("[ %s: %q ]", t.Type, t.Value)
}
