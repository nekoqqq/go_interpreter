package repl

import (
	"bufio"
	"fmt"
	"go_interpreter/lexer"
	"go_interpreter/token"
	"io"
	"log"
)

const PROMPT = ">>"

func Serve(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprintf(out, PROMPT)
		if err != nil {
			log.Printf("输出promt到客户端失败, err: %v\n", err)
			continue
		}
		scanned := scanner.Scan()
		if !scanned {
			log.Printf("scanner失败, err: %v\n", err)
			continue
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", t)
			if err != nil {
				log.Printf("获取token失败, t: %v, err: %v\n", t, err)
				continue
			}
		}
	}
}
