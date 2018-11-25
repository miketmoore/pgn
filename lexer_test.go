package pgn_test

import (
	"fmt"
	"testing"

	pgn "github.com/miketmoore/pgn-3"
)

func TestIsTagPair(t *testing.T) {
	data := []struct {
		name string
		in   string
		out  []pgn.Token
	}{
		{
			name: "Tag Pair",
			in:   "[Event \"F/S Return Match\"]",
			out: []pgn.Token{
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Event"),
				pgn.NewToken(pgn.String, "F/S Return Match"),
				pgn.NewToken(pgn.TagPairClose, "]"),
			},
		},
		{
			name: "Tag Pair - Lots of Whitespace",
			in:   "  [   Event \"F/S      Return      Match\"    ]    ",
			out: []pgn.Token{
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Event"),
				pgn.NewToken(pgn.String, "F/S Return Match"),
				pgn.NewToken(pgn.TagPairClose, "]"),
			},
		},
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			scanner := pgn.NewScanner(test.in)
			lexer := pgn.NewLexer(scanner)

			err, tokens := lexer.Tokenize()

			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Tokens: ", tokens)
				t.Fatal("Unexpected error returned")
			}
			if len(tokens) != len(test.out) {
				fmt.Println(tokens)
				t.Fatal("Unexpected total tokens")
			}
		})
	}
}
