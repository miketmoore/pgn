package pgn_test

import (
	"fmt"
	"testing"

	pgn "github.com/miketmoore/pgn-3"
)

func TestTokenize(t *testing.T) {
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
			in: "[Event \"F/S Return Match\"]\n" +
				"[Site \"Belgrade, Serbia JUG\"]\n" +
				"[Date \"1992.11.04\"]\n[Round \"29\"]\n" +
				"[White \"Fischer, Robert J.\"]\n" +
				"[Black \"Spassky, Boris V.\"]\n" +
				"[Result \"1/2-1/2\"]\n" +
				"[a \"\"]\n" +
				"[A \"\"]\n" +
				"[_ \"\"]\n",
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
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "a"},
				pgn.Token{Type: pgn.String, Value: ""},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "A"},
				pgn.Token{Type: pgn.String, Value: ""},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "_"},
				pgn.Token{Type: pgn.String, Value: ""},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
			},
		},
		{
			name: "Movetext - One",
			in:   "1. e4 e5",
			out: []pgn.Token{
				pgn.Token{Type: pgn.MoveNumber, Value: "1"},
				pgn.Token{Type: pgn.File, Value: "e"},
				pgn.Token{Type: pgn.Rank, Value: "4"},
				pgn.Token{Type: pgn.File, Value: "e"},
				pgn.Token{Type: pgn.Rank, Value: "5"},
			},
		},
		{
			name: "Movetext - Two",
			in:   "1. e4 e5 2. Nf3 Nc6",
			out: []pgn.Token{
				pgn.Token{Type: pgn.MoveNumber, Value: "1"},
				pgn.Token{Type: pgn.File, Value: "e"},
				pgn.Token{Type: pgn.Rank, Value: "4"},
				pgn.Token{Type: pgn.File, Value: "e"},
				pgn.Token{Type: pgn.Rank, Value: "5"},
				pgn.Token{Type: pgn.MoveNumber, Value: "2"},
				pgn.Token{Type: pgn.Piece, Value: "N"},
				pgn.Token{Type: pgn.File, Value: "f"},
				pgn.Token{Type: pgn.Rank, Value: "3"},
				pgn.Token{Type: pgn.Piece, Value: "N"},
				pgn.Token{Type: pgn.File, Value: "c"},
				pgn.Token{Type: pgn.Rank, Value: "6"},
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
