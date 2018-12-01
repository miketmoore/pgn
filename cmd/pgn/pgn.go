package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/miketmoore/pgn"
)

func main() {

	file := flag.TokenString("input", "", "Path to a *.pgn file containing zero or more games")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", *file)
		os.Exit(1)
	}

	scanner := pgn.NewScanner(string(bytes))
	lexer := pgn.NewLexer(scanner)

	tokens := []pgn.Token{}
	err, tokens = lexer.Tokenize(tokens)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Total tokens parsed: ", len(tokens))

	os.Exit(0)
}
