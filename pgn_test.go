package pgn_test

import (
	"fmt"
	"testing"

	. "github.com/miketmoore/pgn"
)

var raw = `[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."]
[Result "1/2-1/2"]

1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.} 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5 Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5 hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6 Nf2 42. g4 Bd3 43. Re6 1/2-1/2`

var parsed = PGN{
	TagPairs: TagPairs{
		Event:  "F/S Return Match",
		Site:   "Belgrade, Serbia JUG",
		Date:   "1992.11.04",
		Round:  "29",
		White:  "Fischer, Robert J.",
		Black:  "Spassky, Boris V.",
		Result: "1/2-1/2",
	},
	Movetext: Movetext{
		MovetextEntry{
			White: Move{Original: "e4", Piece: PiecePawn, File: FileE, Rank: Rank4},
			Black: Move{Original: "e5", Piece: PiecePawn, File: FileE, Rank: Rank5},
		},
		MovetextEntry{
			White: Move{Original: "Nf3", Piece: PieceKnight, File: FileF, Rank: Rank3},
			Black: Move{Original: "Nc6", Piece: PieceKnight, File: FileC, Rank: Rank6},
		},
		MovetextEntry{
			White:    Move{Original: "Bb5", Piece: PieceBishop, File: FileB, Rank: Rank5},
			Black:    Move{Original: "a6", Piece: PiecePawn, File: FileA, Rank: Rank6},
			Comments: []Comment{"This opening is called the Ruy Lopez."},
		},
		MovetextEntry{
			White: Move{Original: "Ba4", File: FileA, Rank: Rank4, Piece: PieceBishop},
			Black: Move{Original: "Nf6", File: FileF, Rank: Rank6, Piece: PieceKnight},
		},
		MovetextEntry{
			White: Move{Original: "O-O", Piece: PieceKing},
			Black: Move{Original: "Be7", File: FileE, Rank: Rank7, Piece: PieceBishop},
		},
		MovetextEntry{
			White: Move{Original: "Re1", File: FileE, Rank: Rank1, Piece: PieceRook},
			Black: Move{Original: "b5", File: FileB, Rank: Rank5, Piece: PiecePawn},
		},
		MovetextEntry{
			White: Move{Original: "Bb3", File: FileB, Rank: Rank3, Piece: PieceBishop},
			Black: Move{Original: "d6", File: FileD, Rank: Rank6, Piece: PiecePawn},
		},
		MovetextEntry{
			White: Move{Original: "c3", File: FileC, Rank: Rank3, Piece: PiecePawn},
			Black: Move{Original: "O-O", Piece: PieceKing},
		},
		MovetextEntry{
			White: Move{Original: "h3", File: FileH, Rank: Rank3, Piece: PiecePawn},
			Black: Move{Original: "Nb8", File: FileB, Rank: Rank8, Piece: PieceKnight},
		},
		MovetextEntry{
			White: Move{Original: "d4", Piece: PiecePawn},
			Black: Move{Original: "Nbd7", Piece: PieceKnight},
		},
		MovetextEntry{
			White: Move{Original: "c4", File: FileC, Rank: Rank4, Piece: PiecePawn},
			Black: Move{Original: "c6", File: FileC, Rank: Rank6, Piece: PiecePawn},
		},
		MovetextEntry{
			White: Move{Original: "cxb5", File: FileB, Rank: Rank5, Piece: PiecePawn, Capture: true},
			Black: Move{Original: "axb5", File: FileB, Rank: Rank5, Piece: PiecePawn, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Nc3", File: FileC, Rank: Rank3, Piece: PieceKnight},
			Black: Move{Original: "Bb7", File: FileB, Rank: Rank7, Piece: PieceBishop},
		},
		MovetextEntry{
			White: Move{Original: "Bg5", File: FileG, Rank: Rank5, Piece: PieceBishop},
			Black: Move{Original: "b4", File: FileB, Rank: Rank4, Piece: PiecePawn},
		},
		MovetextEntry{
			White: Move{Original: "Nb1", File: FileB, Rank: Rank1, Piece: PieceKnight},
			Black: Move{Original: "h6", File: FileH, Rank: Rank6, Piece: PiecePawn},
		},
		MovetextEntry{
			White: Move{Original: "Bh4", File: FileH, Rank: Rank4, Piece: PieceBishop},
			Black: Move{Original: "c5", File: FileC, Rank: Rank5, Piece: PiecePawn},
		},
		MovetextEntry{
			White: Move{Original: "dxe5", File: FileE, Rank: Rank5, Piece: PiecePawn, Capture: true},
			Black: Move{Original: "Nxe4", Piece: PieceKnight, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Bxe7", File: FileE, Rank: Rank7, Piece: PieceBishop, Capture: true},
			Black: Move{Original: "Qxe7", File: FileE, Rank: Rank7, Piece: PieceQueen, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "exd6", File: FileD, Rank: Rank6, Piece: PiecePawn, Capture: true},
			Black: Move{Original: "Qf6", File: FileF, Rank: Rank6, Piece: PieceQueen, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Nbd2", File: FileD, Rank: Rank2, Piece: PieceKnight},
			Black: Move{Original: "Nxd6", File: FileD, Rank: Rank6, Piece: PieceKnight, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Nc4", File: FileC, Rank: Rank4, Piece: PieceKnight},
			Black: Move{Original: "Nxc4", File: FileC, Rank: Rank4, Piece: PieceKnight, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Bxc4", File: FileC, Rank: Rank4, Piece: PieceBishop, Capture: true},
			Black: Move{Original: "Nb6", File: FileB, Rank: Rank6, Piece: PieceKnight},
		},
		MovetextEntry{
			White: Move{Original: "Ne5", File: FileE, Rank: Rank5, Piece: PieceKnight},
			Black: Move{Original: "Rae8", File: FileE, Rank: Rank8, Piece: PieceRook},
		},
		MovetextEntry{
			White: Move{Original: "Bxf7+", File: FileF, Rank: Rank7, Piece: PieceBishop, Capture: true, Check: true},
			Black: Move{Original: "Rxf7", File: FileF, Rank: Rank7, Piece: PieceRook, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Nxf7", File: FileF, Rank: Rank7, Piece: PieceKnight, Capture: true},
			Black: Move{Original: "Rxe1+", File: FileE, Rank: Rank1, Piece: PieceRook, Capture: true, Check: true},
		},
		MovetextEntry{
			White: Move{Original: "Qxe1", File: FileE, Rank: Rank1, Piece: PieceQueen, Capture: true},
			Black: Move{Original: "Kxf7", File: FileF, Rank: Rank7, Piece: PieceKing, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Qe3", File: FileE, Rank: Rank3, Piece: PieceQueen},
			Black: Move{Original: "Qg5", File: FileG, Rank: Rank5, Piece: PieceQueen},
		},
		MovetextEntry{
			White: Move{Original: "Qxg5", File: FileG, Rank: Rank5, Piece: PieceQueen, Capture: true},
			Black: Move{Original: "hxg5", File: FileG, Rank: Rank5, Piece: PiecePawn, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "b3", File: FileB, Rank: Rank3, Piece: PiecePawn},
			Black: Move{Original: "Ke6", File: FileE, Rank: Rank6, Piece: PieceKing},
		},
		MovetextEntry{
			White: Move{Original: "a3", File: FileA, Rank: Rank3, Piece: PiecePawn},
			Black: Move{Original: "Kd6", File: FileD, Rank: Rank6, Piece: PieceKing},
		},
		MovetextEntry{
			White: Move{Original: "axb4", File: FileB, Rank: Rank4, Piece: PiecePawn, Capture: true},
			Black: Move{Original: "cxb4", File: FileB, Rank: Rank4, Piece: PiecePawn, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Ra5", File: FileA, Rank: Rank5, Piece: PieceRook},
			Black: Move{Original: "Nd5", File: FileD, Rank: Rank5, Piece: PieceKnight},
		},
		MovetextEntry{
			White: Move{Original: "f3", File: FileF, Rank: Rank3, Piece: PiecePawn},
			Black: Move{Original: "Bc8", File: FileC, Rank: Rank8, Piece: PieceBishop},
		},
		MovetextEntry{
			White: Move{Original: "Kf2", File: FileF, Rank: Rank2, Piece: PieceKing},
			Black: Move{Original: "Bf5", File: FileF, Rank: Rank5, Piece: PieceBishop},
		},
		MovetextEntry{
			White: Move{Original: "Ra7", File: FileA, Rank: Rank7, Piece: PieceRook},
			Black: Move{Original: "g6", File: FileG, Rank: Rank6, Piece: PiecePawn},
		},
		MovetextEntry{
			White: Move{Original: "Ra6+", File: FileA, Rank: Rank6, Piece: PieceRook, Check: true},
			Black: Move{Original: "Kc5", File: FileC, Rank: Rank5, Piece: PieceKing},
		},
		MovetextEntry{
			White: Move{Original: "Ke1", File: FileE, Rank: Rank1, Piece: PieceKing},
			Black: Move{Original: "Nf4", File: FileF, Rank: Rank4, Piece: PieceKnight},
		},
		MovetextEntry{
			White: Move{Original: "g3", File: FileG, Rank: Rank3, Piece: PiecePawn},
			Black: Move{Original: "Nxh3", File: FileH, Rank: Rank3, Piece: PieceKnight, Capture: true},
		},
		MovetextEntry{
			White: Move{Original: "Kd2", File: FileD, Rank: Rank2, Piece: PieceKing},
			Black: Move{Original: "Kb5", File: FileB, Rank: Rank5, Piece: PieceKing},
		},
		MovetextEntry{
			White: Move{Original: "Rd6", File: FileD, Rank: Rank6, Piece: PieceRook},
			Black: Move{Original: "Kc5", File: FileC, Rank: Rank5, Piece: PieceKing},
		},
		MovetextEntry{
			White: Move{Original: "Ra6", File: FileA, Rank: Rank6, Piece: PieceRook},
			Black: Move{Original: "Nf2", File: FileF, Rank: Rank2, Piece: PieceKnight},
		},
		MovetextEntry{
			White: Move{Original: "g4", File: FileG, Rank: Rank4, Piece: PiecePawn},
			Black: Move{Original: "Bd3", File: FileD, Rank: Rank3, Piece: PieceBishop},
		},
		MovetextEntry{
			White: Move{Original: "Re6", File: FileE, Rank: Rank6, Piece: PieceRook},
			Black: Move{Original: "1/2-1/2"},
		},
	},
}

func TestParse(t *testing.T) {
	got := Parse(raw)
	if got.TagPairs != parsed.TagPairs {
		fmt.Printf("Got:\n%v\n", got)
		fmt.Printf("Expected:\n%v\n", parsed)
		t.Fatal("Parsing PGN string failed")
	}
	if !assertMovetextEquality(parsed.Movetext, got.Movetext) {
		fmt.Printf("Got:\n%v\n", got)
		fmt.Printf("Expected:\n%v\n", parsed)
		t.Fatal("Movetext is unexpected during PGN parsing")
	}
}

func TestOuptut(t *testing.T) {
	got := parsed.String()
	if got != raw {
		fmt.Printf("Got:\n%s\n", got)
		fmt.Printf("Expected:\n%s\n", raw)
		t.Fatal("output failed")
	}
}

func assertMovetextEquality(a, b Movetext) bool {
	if len(a) != len(b) {
		return false
	}
	for i, valA := range a {
		if valA.White.Original != b[i].White.Original || valA.Black.Original != b[i].Black.Original {
			return false
		}
		if len(valA.Comments) != len(b[i].Comments) {
			return false
		}
	}
	return true
}
