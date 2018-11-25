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
		{
			name: "Multiple Tag Pairs",
			/*
				[Event "F/S Return Match"]
				[Site "Belgrade, Serbia JUG"]
				[Date "1992.11.04"]
				[Round "29"]
				\n[White \"Fischer, Robert J.\"]\n[Black \"Spassky, Boris V.\"]\n[Result \"1/2-1/2\"]
			*/
			in: "[Event \"F/S Return Match\"]\n[Site \"Belgrade, Serbia JUG\"]\n[Date \"1992.11.04\"]\n[Round \"29\"]\n[White \"Fischer, Robert J.\"]\n[Black \"Spassky, Boris V.\"]\n[Result \"1/2-1/2\"]",
			out: []pgn.Token{
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Event"),
				pgn.NewToken(pgn.String, "F/S Return Match"),
				pgn.NewToken(pgn.TagPairClose, "]"),
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Site"),
				pgn.NewToken(pgn.String, "Belgrade, Serbia JUG"),
				pgn.NewToken(pgn.TagPairClose, "]"),
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Date"),
				pgn.NewToken(pgn.String, "1992.11.04"),
				pgn.NewToken(pgn.TagPairClose, "]"),
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Round"),
				pgn.NewToken(pgn.String, "29"),
				pgn.NewToken(pgn.TagPairClose, "]"),
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "White"),
				pgn.NewToken(pgn.String, "Fischer, Robert J."),
				pgn.NewToken(pgn.TagPairClose, "]"),
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Black"),
				pgn.NewToken(pgn.String, "Spassky, Boris V."),
				pgn.NewToken(pgn.TagPairClose, "]"),
				pgn.NewToken(pgn.TagPairOpen, "["),
				pgn.NewToken(pgn.TagName, "Result"),
				pgn.NewToken(pgn.String, "1/2-1/2"),
				pgn.NewToken(pgn.TagPairClose, "]"),
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
