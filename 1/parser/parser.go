/* parser is a software component that takes input and builds a data structure, often some kind of parse tree like js object. lexer feeds tokens into parser 
this is going to be a Pratt parser(top-down operator precedence parser) which is a type of recursive descent parser.
bindings in this language are in the for of let, where the <identifier> = <expression> expressions produce values, statements don't
*/

package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)


//parser has three fields: l, curToken and peekToken
//l is a pointer to an instance of a lexer where NextToken() is repeatedly called
type Parser struct {
	l *lexer.Lexer
	
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//read two tokens, so curToken and peekToken are both set
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