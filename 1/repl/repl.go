//REPL (Read Eval Print Loop) in js the repl is 'node ...'
//tokenizes src code and prints tokens
package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "
//read from the input src until we encounter a new lin, take the just read line and pass it to instance of lexer and finally print tokens until eof
func Start(in io.Reader, out io.Writer) {
scanner := bufio.NewScanner(in)	

for {
	fmt.Fprint(out, PROMPT)
	scanned := scanner.Scan()
	if !scanned {
		return
	}

	line := scanner.Text()
	l := lexer.New(line)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { 
		fmt.Fprintf(out, "%+v\n", tok)
	}
}
}