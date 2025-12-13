package parse

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parse struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parse {
	p := &Parse{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parse) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parse) ParseProgram() *ast.Program {
	return nil
}