package token
//by defining token type as string, we can use many different values as TokenTypes, which lets us distinguish btwn different tokens
type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"//token we don't know about 
	EOF = "EOF"//end of file, tells parser later that it can stop

	//Identifiers + literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT = "INT" //123456
	EQ = "=="
	NOT_EQ = "!="

	//OPERATORS
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	//Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

//checks keywords table to see whether the given identifier is in fact a keyword
// if it is, it returns the keyword's TokenType constant
//if not, we just get back token.IDENT which is the type for all user-defined identifiers
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}