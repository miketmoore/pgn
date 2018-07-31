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

	{pgn.CASTLE_QUEEN_SIDE, "O-O-O", true},
	{pgn.CASTLE_QUEEN_SIDE, "O-O", false},

	{pgn.RESULT_DRAW, "1/2-1/2", true},
	{pgn.RESULT_WHITE_WINS, "1-0", true},
	{pgn.RESULT_BLACK_WINS, "0-1", true},
	{pgn.RESULT_UNFINISHED, "*", true},

	{pgn.SINGLE_DIGIT_MOVE_NO, "0.", false},
	{pgn.SINGLE_DIGIT_MOVE_NO, "1.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "2.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "3.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "4.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "5.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "6.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "7.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "8.", true},
	{pgn.SINGLE_DIGIT_MOVE_NO, "9.", true},

	{pgn.DOUBLE_DIGIT_MOVE_NO, "00.", false},
	{pgn.DOUBLE_DIGIT_MOVE_NO, "1.", false},
	{pgn.DOUBLE_DIGIT_MOVE_NO, "10.", true},
	{pgn.DOUBLE_DIGIT_MOVE_NO, "20.", true},
	{pgn.DOUBLE_DIGIT_MOVE_NO, "30.", true},
	{pgn.DOUBLE_DIGIT_MOVE_NO, "40.", true},
	{pgn.DOUBLE_DIGIT_MOVE_NO, "50.", true},
	{pgn.DOUBLE_DIGIT_MOVE_NO, "60.", false},

	{pgn.TAGPAIR_START, "[", true},
	{pgn.TAGPAIR_END, "]", true},

	{pgn.TAG_EVENT, "Event", true},
	{pgn.TAG_SITE, "Site", true},
	{pgn.TAG_DATE, "Date", true},
	{pgn.TAG_ROUND, "Round", true},
	{pgn.TAG_WHITE, "White", true},
	{pgn.TAG_BLACK, "Black", true},
	{pgn.TAG_RESULT, "Result", true},

	{pgn.TAG_VALUE, "\"tag value\"", true},

	{pgn.TAG_VALUE_DATE, "????.??.??", true},
	{pgn.TAG_VALUE_DATE, "2018.07.31", true},
	{pgn.TAG_VALUE_DATE, "2018.07.??", true},
	{pgn.TAG_VALUE_DATE, "2018.??.??", true},
}

func TestTokens(t *testing.T) {
	for i, test := range tests {
		t.Run(string(i), func(t *testing.T) {
			match := test.re.MatchString(test.in)
			if match != test.match {
				fmt.Printf("Regexp: %v\n", test.re)
				fmt.Printf("Input: %s\n", test.in)
				fmt.Println("Expected match: ", test.match)
				fmt.Println("Got match: ", match)
				t.Fatal("Token regexp failed")
			}
		})
	}
}
