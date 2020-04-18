//parser/parser.go

package parser


import (
	"github.com/Rasino66/Monkey/ast"
	"github.com/Rasino66/Monkey/lexer"
	"github.com/Rasino66/Monkey/token"
)

type Parser struct {
  l *lexer.Lexer

  curToken token.Token
  peekToken token.Token
}


func New(l *lexer.Lexer) *Parser {
  p := &Parser{l: l}

  // 2 token read , set curToken peekToken
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
