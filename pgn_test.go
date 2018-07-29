package pgn_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/pgn"
)

// var game = `[Event "F/S Return Match"]
// [Site "Belgrade, Serbia JUG"]
// [Date "1992.11.04"]
// [Round "29"]
// [White "Fischer, Robert J."]
// [Black "Spassky, Boris V."]
// [Result "1/2-1/2"]

// 1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}
// 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7
// 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5
// Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6
// 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5
// hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5
// 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6
// Nf2 42. g4 Bd3 43. Re6 1/2-1/2`

var raw = `[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."]
[Result "1/2-1/2"]

1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.} 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5 Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5 hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6 Nf2 42. g4 Bd3 43. Re6 1/2-1/2`

var parsed = pgn.PGN{
	TagPairs: pgn.TagPairs{
		Event:  "F/S Return Match",
		Site:   "Belgrade, Serbia JUG",
		Date:   "1992.11.04",
		Round:  "29",
		White:  "Fischer, Robert J.",
		Black:  "Spassky, Boris V.",
		Result: "1/2-1/2",
	},
	Movetext: pgn.Movetext{
		pgn.MovetextEntry{White: "e4", Black: "e5"},
		pgn.MovetextEntry{White: "Nf3", Black: "Nc6"},
		pgn.MovetextEntry{White: "Bb5", Black: "a6", Comments: []pgn.Comment{"This opening is called the Ruy Lopez."}},
		pgn.MovetextEntry{White: "Ba4", Black: "Nf6"},
		pgn.MovetextEntry{White: "O-O", Black: "Be7"},
		pgn.MovetextEntry{White: "Re1", Black: "b5"},
		pgn.MovetextEntry{White: "Bb3", Black: "d6"},
		pgn.MovetextEntry{White: "c3", Black: "O-O"},
		pgn.MovetextEntry{White: "h3", Black: "Nb8"},
		pgn.MovetextEntry{White: "d4", Black: "Nbd7"},
		pgn.MovetextEntry{White: "c4", Black: "c6"},
		pgn.MovetextEntry{White: "cxb5", Black: "axb5"},
		pgn.MovetextEntry{White: "Nc3", Black: "Bb7"},
		pgn.MovetextEntry{White: "Bg5", Black: "b4"},
		pgn.MovetextEntry{White: "Nb1", Black: "h6"},
		pgn.MovetextEntry{White: "Bh4", Black: "c5"},
		pgn.MovetextEntry{White: "dxe5", Black: "Nxe4"},
		pgn.MovetextEntry{White: "Bxe7", Black: "Qxe7"},
		pgn.MovetextEntry{White: "exd6", Black: "Qf6"},
		pgn.MovetextEntry{White: "Nbd2", Black: "Nxd6"},
		pgn.MovetextEntry{White: "Nc4", Black: "Nxc4"},
		pgn.MovetextEntry{White: "Bxc4", Black: "Nb6"},
		pgn.MovetextEntry{White: "Ne5", Black: "Rae8"},
		pgn.MovetextEntry{White: "Bxf7+", Black: "Rxf7"},
		pgn.MovetextEntry{White: "Nxf7", Black: "Rxe1+"},
		pgn.MovetextEntry{White: "Qxe1", Black: "Kxf7"},
		pgn.MovetextEntry{White: "Qe3", Black: "Qg5"},
		pgn.MovetextEntry{White: "Qxg5", Black: "hxg5"},
		pgn.MovetextEntry{White: "b3", Black: "Ke6"},
		pgn.MovetextEntry{White: "a3", Black: "Kd6"},
		pgn.MovetextEntry{White: "axb4", Black: "cxb4"},
		pgn.MovetextEntry{White: "Ra5", Black: "Nd5"},
		pgn.MovetextEntry{White: "f3", Black: "Bc8"},
		pgn.MovetextEntry{White: "Kf2", Black: "Bf5"},
		pgn.MovetextEntry{White: "Ra7", Black: "g6"},
		pgn.MovetextEntry{White: "Ra6+", Black: "Kc5"},
		pgn.MovetextEntry{White: "Ke1", Black: "Nf4"},
		pgn.MovetextEntry{White: "g3", Black: "Nxh3"},
		pgn.MovetextEntry{White: "Kd2", Black: "Kb5"},
		pgn.MovetextEntry{White: "Rd6", Black: "Kc5"},
		pgn.MovetextEntry{White: "Ra6", Black: "Nf2"},
		pgn.MovetextEntry{White: "g4", Black: "Bd3"},
		pgn.MovetextEntry{White: "Re6", Black: "1/2-1/2"},
	},
}

func TestParse(t *testing.T) {
	got := pgn.Parse(raw)
	if got.TagPairs != parsed.TagPairs {
		fmt.Printf("Got:\n%s\n", got)
		fmt.Printf("Expected:\n%s\n", parsed)
		t.Fatal("Parsing PGN string failed")
	}
	if !assertMovetextEquality(parsed.Movetext, got.Movetext) {
		fmt.Printf("Got:\n%s\n", got)
		fmt.Printf("Expected:\n%s\n", parsed)
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

func assertMovetextEquality(a, b pgn.Movetext) bool {
	if len(a) != len(b) {
		return false
	}
	for i, valA := range a {
		if valA.White != b[i].White || valA.Black != b[i].Black {
			return false
		}
		if len(valA.Comments) != len(b[i].Comments) {
			return false
		}
	}
	return true
}
