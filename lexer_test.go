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
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Event"},
				pgn.Token{Type: pgn.String, Value: "F/S Return Match"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
			},
		},
		{
			name: "Tag Pair - Lots of Whitespace",
			in:   "  [   Event \"F/S      Return      Match\"    ]    ",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Event"},
				pgn.Token{Type: pgn.String, Value: "F/S Return Match"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
			},
		},
		{
			name: "Multiple Tag Pairs",
			in:   "[Event \"F/S Return Match\"]\n[Site \"Belgrade, Serbia JUG\"]\n[Date \"1992.11.04\"]\n[Round \"29\"]\n[White \"Fischer, Robert J.\"]\n[Black \"Spassky, Boris V.\"]\n[Result \"1/2-1/2\"]",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Event"},
				pgn.Token{Type: pgn.String, Value: "F/S Return Match"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Site"},
				pgn.Token{Type: pgn.String, Value: "Belgrade, Serbia JUG"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Date"},
				pgn.Token{Type: pgn.String, Value: "1992.11.04"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Round"},
				pgn.Token{Type: pgn.String, Value: "29"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "White"},
				pgn.Token{Type: pgn.String, Value: "Fischer, Robert J."},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Black"},
				pgn.Token{Type: pgn.String, Value: "Spassky, Boris V."},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Result"},
				pgn.Token{Type: pgn.String, Value: "1/2-1/2"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
			},
		},
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			scanner := pgn.NewScanner(test.in)
			lexer := pgn.NewLexer(scanner)

			tokens := []pgn.Token{}
			err, tokens := lexer.Tokenize(tokens)
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
