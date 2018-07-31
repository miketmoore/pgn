package pgn_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/miketmoore/pgn"
)

var tests = []struct {
	re    *regexp.Regexp
	in    string
	match bool
}{
	{pgn.COMMENT, "", false},
	{pgn.COMMENT, "    ", false},
	{pgn.COMMENT, "{}", true},
	{pgn.COMMENT, "{comment}", true},
	{pgn.COMMENT, "{ here is a comment }", true},

	{pgn.PAWN_MOVE, "a1", true},
	{pgn.PAWN_MOVE, "Bh8", false},

	{pgn.NON_PAWN_MOVE, "Ba1", true},
	{pgn.NON_PAWN_MOVE, "a1", false},

	{pgn.PAWN_CAPTURE, "bxc2", true},
	{pgn.PAWN_CAPTURE, "Bxc2", false},

	{pgn.NON_PAWN_CAPTURE, "Bxc2", true},
	{pgn.NON_PAWN_CAPTURE, "bxc2", false},

	{pgn.CASTLE_KING_SIDE, "O-O", true},
	{pgn.CASTLE_KING_SIDE, "O-O-O", false},
}

func TestTokens(t *testing.T) {
	for i, test := range tests {
		t.Run(string(i), func(t *testing.T) {
			match := test.re.MatchString(test.in)
			if match != test.match {
				fmt.Printf("%v\n", test.re)
				t.Fatal("failed: ", match)
			}
		})
	}
}
