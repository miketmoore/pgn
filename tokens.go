package pgn

import (
	"regexp"
)

var (
	COMMENT              = regexp.MustCompile(`{.*}`)
	PAWN_MOVE            = regexp.MustCompile(`^[a-h][1-8]$`)
	NON_PAWN_MOVE        = regexp.MustCompile(`[NBRQK][a-h][1-8]`)
	PAWN_CAPTURE         = regexp.MustCompile(`[a-h]x[a-h][1-8]`)
	NON_PAWN_CAPTURE     = regexp.MustCompile(`[NBRQK]x[a-h][1-8]`)
	CASTLE_KING_SIDE     = regexp.MustCompile(`^O-O$`)
	CASTLE_QUEEN_SIDE    = regexp.MustCompile(`O-O-O`)
	RESULT_DRAW          = regexp.MustCompile(`1/2-1/2`)
	RESULT_WHITE_WINS    = regexp.MustCompile(`1-0`)
	RESULT_BLACK_WINS    = regexp.MustCompile(`0-1`)
	RESULT_UNFINISHED    = regexp.MustCompile(`\*`)
	SINGLE_DIGIT_MOVE_NO = regexp.MustCompile(`[1-9]{1}\.`)
	DOUBLE_DIGIT_MOVE_NO = regexp.MustCompile(`[1-5]{1}[0-9]{1}\.`)
	TAGPAIR_START        = regexp.MustCompile(`\[`)
	TAGPAIR_END          = regexp.MustCompile(`\]`)
	TAG_EVENT            = regexp.MustCompile(`Event`)
	TAG_SITE             = regexp.MustCompile(`Site`)
	TAG_DATE             = regexp.MustCompile(`Date`)
	TAG_ROUND            = regexp.MustCompile(`Round`)
	TAG_WHITE            = regexp.MustCompile(`White`)
	TAG_BLACK            = regexp.MustCompile(`Black`)
	TAG_RESULT           = regexp.MustCompile(`Result`)
	TAG_VALUE            = regexp.MustCompile(`\"(.*)\"`)
	TAG_VALUE_DATE       = regexp.MustCompile(`[0-9?]{4}\.[0-9?]{2}\.[0-9?]{2}`)
)
