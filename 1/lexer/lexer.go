//Goal of this is to take source code as input and output the tokens that rep src code. It does not need to buffer or save tokens since there is only one method: NextToken()
//If we were to attach line & col numbers to tokens we would initialize the lexer with an io.Reader and the filename. Since that is more complex we'll just stick to TYPE: STRING
package lexer

import "monkey/token"

type Lexer struct {
	input 			 string
	position		 int //current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch					 byte	//current char nder examination
}

//ex: l.input[l.readPosition] with l being the char
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
		
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-': 
		tok = newToken(token.MINUS, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch //saveing to local var before readchar is called again so we dont lose the value of the current char
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		//use readIdentifier to set the Literal field  of our current token
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
}

l.readChar()
return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//similar to readchar but doesn't advance position of l.position and l.readPosition
//used to look ahead in input and see what the next char is without advancing the lexer
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//gives next char and advances position in the input string
//if it reaches the end of an input, if it has, it sets l.ch to 0 (aka null)
//if not, it advances l.ch to the next char by accessing l.input[l.readPosition] -> it's then updated to to the just used l.readPosition and is incremented by 1
//supports ASCII instead of unicode for simplicity
//if we were to do that, we would need to change l.ch from byte to rune and change the way we read the next chars since they could be multiple byte wide now
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}