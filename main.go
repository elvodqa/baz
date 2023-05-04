package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/elvodqa/baz/compiler/lexer"
)

var (
	SourcePath string = "."
	SourceFile string = ""
	ReplMode   string = "lexer"
)

func init() {
	fmt.Println("Baz v0.0.1")
	fmt.Println("------------------")
	flag.StringVar(&SourcePath, "src", "", "Source path")
	flag.StringVar(&SourceFile, "file", "", "Source file")
	flag.StringVar(&ReplMode, "repl", "", "Lexer REPL mode")
	flag.Parse()
}

func main() {
	switch ReplMode {
	case "lexer":
		fmt.Println("Lexer REPL mode")

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')
			if text == "##exit\n" || text == "##quit\n" {
				return
			}
			if text == "##clear\n" {
				fmt.Print("\033[H\033[2J")
				continue
			}
			l := lexer.NewLexer(text)
			for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		}
	case "parser":
		break
	}
}
