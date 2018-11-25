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
			name: "Letters",
			in:   "[Event \"F/S Return Match\"]",
			out: []pgn.Token{
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Event"),
				pgn.NewToken(pgn.TagValue, "F/S Return Match"),
				pgn.NewToken(pgn.TagPairClose, "]"),
			},
		},
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			scanner := pgn.NewScanner(test.in)
			lexer := pgn.NewLexer(scanner)

			ok, tokens := lexer.Tokenize()

			fmt.Println(ok)
			fmt.Println(tokens)
			// if !ok {
			// 	fmt.Println(tokens)
			// 	t.Fatal("Not OK")
			// }
			// if len(tokens) != len(test.out) {
			// 	t.Fatal("failed")
			// }
		})
	}
}
