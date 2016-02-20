package parser

import (
	"fmt"

	"github.com/bbuck/dragonscript/ast"
	"github.com/bbuck/dragonscript/lexer"
)

// Parser represents an object that will consume a list of tokens and produce an
// Abstract Syntax Tree that can then be run through an Interpreter.
type Parser struct {
	line   int
	tokens []*lexer.Token
	pos    int
	ast    ast.Node
	err    error
}

func NewParser(ts []*lexer.Token) *Parser {
	return &Parser{
		tokens: ts,
		pos:    -1,
	}
}

func (p *Parser) Parse() (ast.Node, error) {
	if p.ast == nil {
		p.ast, p.err = p.parseExpressionList()
	}

	return p.ast, p.err
}

func (p *Parser) expect(t lexer.TokenType) (*lexer.Token, error) {
	next := p.nextToken()
	if next.Type == t {
		return next, nil
	}

	return nil, unexpectedTokenError(next)
}

func (p *Parser) nextToken() *lexer.Token {
	if p.pos < len(p.tokens)-1 {
		p.pos++
		return p.tokens[p.pos]
	}

	return nil
}

func (p *Parser) rewind() {
	p.pos--

	if p.pos < 0 {
		p.pos = 0
	}
}

func (p *Parser) peek(n int) *lexer.Token {
	var t *lexer.Token
	for i := 0; i < n; i++ {
		t = p.nextToken()
	}

	for i := 0; i < n; i++ {
		p.rewind()
	}

	return t
}

func unexpectedTokenError(t *lexer.Token) error {
	return fmt.Errorf("unexpected token %s encountered", t.Type)
}
