package parser

import (
	"github.com/SeaSkyThe/MonkeyInterpreter/ast"
	"github.com/SeaSkyThe/MonkeyInterpreter/lexer"
	"github.com/SeaSkyThe/MonkeyInterpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, to set curToken and peekToken

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
