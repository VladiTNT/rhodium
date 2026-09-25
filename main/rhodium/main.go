package main

import (
	"bufio"
	"fmt"
	"os"
	"rhodium/src/interpreter"
	"rhodium/src/lexer"
	"rhodium/src/parser"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print(">> ")
	for sc.Scan() {
		line := sc.Text()
		if line == "exit" {
			break
		}

		tokens := lexer.Tokenize(strings.NewReader(line))
		tree, _ := parser.PrattParse(tokens, 0, 0)
		fmt.Println(interpreter.Do(tree))

		fmt.Print(">> ")
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}
}
